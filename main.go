package main

type global struct {

}

func main(){
	r := setupRouter()
	r.Run(":7070")
}
