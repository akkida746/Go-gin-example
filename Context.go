package main
import ("fmt";"context")

//https://levelup.gitconnected.com/context-in-golang-98908f042a57
func main(){
	ctx := context.Background()
	context.WithValue(ctx, "a", "test-value")

	fmt.Println(ctx.Value("a"))
}