package main
import("bytes";"encoding/json";"flag";"fmt";"math/rand";"net/http";"github.com/Mairaelefante/go-collision-anomaly-detector/internal/core")
func main(){n:=flag.Int("count",1000,"events");flag.Parse();r:=rand.New(rand.NewSource(42));for i:=0;i<*n;i++{a:=r.Float64()<.04;s:=1.;if a{s=4};e:=core.Event{EventID:fmt.Sprintf("evt-%06d",i),TotalPT:mathmax(0,150+r.NormFloat64()*55)*s,MissingET:mathmax(0,35+r.NormFloat64()*18)*s,JetCount:int(mathmax(0,2.4+r.NormFloat64()*1.2*s)),LeptonCount:int(mathmax(0,.8+r.NormFloat64()*.8*s)),InvariantMass:mathmax(0,91.2+r.NormFloat64()*16*s),DetectorOccupancy:mathmin(1,mathmax(0,.4+r.NormFloat64()*.13*s))};b,_:=json.Marshal(e);resp,err:=http.Post("http://localhost:8080/events","application/json",bytes.NewReader(b));if err!=nil{panic(err)};resp.Body.Close()};fmt.Println("submitted",*n)}
func mathmax(a,b float64)float64{if a>b{return a};return b};func mathmin(a,b float64)float64{if a<b{return a};return b}

