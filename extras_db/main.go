package extrasdb

import (
	"context"
	"errors"
	"fmt"
)

func Wrapper(ctx context.Context, callback func(ctx context.Context) error) error {
	ctxValue := ctx.Value("key")
	fmt.Println("Context Value is: ", ctxValue)
	if ctxValue.(int)%2 != 0 {
		fmt.Println("ERROR WITH VALUE FROM CONTEXT")
		return errors.New("ERROR WITH CONTEXT VALUE")
	}

	fmt.Println("ALL GOOD WITH CONTEXT VALUE. NOW, LET's VALIDATE")
	if err := callback(ctx); err != nil {
		fmt.Println("ERROR EN LOGICA:", err.Error())
	}
	return nil
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", 3)
	err := Wrapper(ctx, func(ctx context.Context) error {
		fmt.Println("LLAMADA DEL CALLBACK ACTIVADA")

		ctxValue := ctx.Value("key")
		if ctxValue.(int) == 4 {
			return errors.New("ERROR VALUE IS EVEN BUT IS 4")
		}
		fmt.Println("FUNCION COMPLETADA")
		return nil
	})
	if err != nil {
		fmt.Println("ERROR DE LLAMADA: ", err.Error())
	}
}
