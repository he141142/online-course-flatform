package errgrouppoc

import (
	"context"
	_ "crypto/sha1"
	"errors"
	"fmt"
	_ "path"
	_ "strings"

	"golang.org/x/sync/errgroup"
)

type CarModel struct {
	Name        string
	Manufacture string
}

func Run(approach string) {

	ctxBackgr := context.Background()
	//ctxBackgr, cancelFn := context.WithCancel(ctxBackgr)
	g, parentCtx := errgroup.WithContext(ctxBackgr)
	//TODO: IMPLEMENT RECEIVERS

	// g.Go(func() error {
	// 	<-ctx.Done()
	// 	cancelFn()
	// 	return nil
	// })
	//

	cars := []CarModel{
		{
			Name:        "BENTLEY CONTINENTAL GT V8 COUPE 2017",
			Manufacture: "Bentley",
		},

		{
			Name:        "BENTLEY CONTINENTAL FLYING SPUR SEDAN 2016",
			Manufacture: "Bentley",
		},
		{
			Name:        "BENTLEY CONTINENTAL GTC 100TH ANNIVERSARY CONVERTIBLE 2023",
			Manufacture: "Bentley",
		}, {
			Name:        "BMW 2-SERIES 218I MINI MPV 2015",
			Manufacture: "BMW",
		}, {

			Name:        "BMW 1 SERIES SPORT HATCHBACK 2022",
			Manufacture: "BMW",
		}, {
			Name:        "BMW 3-SERIES 328I XDRIVE GRAN TURISMO HATCHBACK 2015",
			Manufacture: "BMW",
		}, {
			Name:        "HONDA ACCORD EX SEDAN 2023",
			Manufacture: "Honda",
		}, {
			Name:        "HONDA ACCORD EX-L SEDAN 2017",
			Manufacture: "Honda",
		},
	}

	switch approach {
	case "1st":
		ImplementStrategy1st(parentCtx, cars, g)
	case "2nd":
		Implement2ndStrategy(parentCtx, cars, g)
	case "3rd":
		ImplementStartegy3(parentCtx, cars, g)
	default:
		panic(errors.New("error no implemented"))

	}

	g.Wait()
}

func ImplementStrategy1st(ctx context.Context, cars []CarModel, g *errgroup.Group) {
	hondaChan := make(chan string)
	bmwChan := make(chan string)
	bentleyChan := make(chan string)

	g.Go(func() error {
		defer func() {
			close(hondaChan)
			close(bmwChan)
			close(bentleyChan)
		}()
		for _, car := range cars {
			switch car.Manufacture {
			case "Honda":
				hondaChan <- car.Name
			case "BMW":
				bmwChan <- car.Name
			case "Bentley":
				bentleyChan <- car.Name
			}
		}

		return nil
	})

	g.Go(func() error {
		subGroup, _ := errgroup.WithContext(ctx)

		subGroup.Go(func() error {
			for c := range bentleyChan {
				fmt.Println(c)
			}
			return nil
		})

		subGroup.Go(func() error {
			for c := range bmwChan {
				fmt.Println(c)
			}
			return nil
		})
		subGroup.Go(func() error {
			for c := range hondaChan {
				fmt.Println(c)
			}
			return nil
		})

		return subGroup.Wait()
	})

}

func Implement2ndStrategy(ctx context.Context, cars []CarModel, g *errgroup.Group) {
	mapChannel := make(map[string]chan string)

	for _, car := range cars {
		if _, exist := mapChannel[car.Manufacture]; exist {
			continue
		}
		mapChannel[car.Manufacture] = make(chan string)
	}

	g.Go(func() error {
		defer func() {
			for _, logChan := range mapChannel {
				close(logChan)
			}
		}()
		for _, car := range cars {
			logCh := mapChannel[car.Manufacture]
			logCh <- car.Name
		}
		return nil
	})

	for _, logCh := range mapChannel {
		//capture curr value
		logCh := logCh
		g.Go(func() error {
			for data := range logCh {
				fmt.Println(data)
			}
			return nil
		})
	}

}

func ImplementStartegy3(ctx context.Context, cars []CarModel, g *errgroup.Group) {
	mapChannel := make(map[string]chan string)

	done := make(chan interface{})

	for _, car := range cars {
		if _, exist := mapChannel[car.Manufacture]; exist {
			continue
		}
		mapChannel[car.Manufacture] = make(chan string)
	}

	g.Go(func() error {
		defer func() {
			for _, logChan := range mapChannel {
				close(logChan)
			}
		}()
		for _, car := range cars {
			logCh := mapChannel[car.Manufacture]
			if car.Name == "BMW 3-SERIES 328I XDRIVE GRAN TURISMO HATCHBACK 2015" {
				close(done)
				return nil
			}
			logCh <- car.Name

		}
		return nil
	})

	for _, logCh := range mapChannel {
		//capture curr value
		logCh := logCh
		g.Go(func() error {
			for {
				select {
				case data, ok := <-logCh:
					if !ok {
						return nil
					}
					fmt.Println(data)
				case <-done:
					fmt.Println("Completed")
					return nil
				}
			}
		})
	}

}
