package core
import("math";"sort";"sync")
type Event struct{EventID string `json:"event_id"`; TotalPT,MissingET float64 `json:"total_pt"`; JetCount,LeptonCount int `json:"jet_count"`; InvariantMass,DetectorOccupancy float64 `json:"invariant_mass"`}
type Result struct{Event Event `json:"event"`; Score float64 `json:"score"`; IsAnomaly bool `json:"is_anomaly"`; Reasons []string `json:"reasons"`}
type Detector struct{mu sync.RWMutex; received,processed uint64; flagged []Result}
func New()*Detector{return &Detector{flagged:make([]Result,0)}}
func(d *Detector)Evaluate(e Event)Result{names:=[]string{"total_pt","missing_et","jet_count","lepton_count","invariant_mass","occupancy"};v:=[]float64{e.TotalPT,e.MissingET,float64(e.JetCount),float64(e.LeptonCount),e.InvariantMass,e.DetectorOccupancy};m:=[]float64{150,35,2.4,.8,91.2,.40};s:=[]float64{55,18,1.2,.8,16,.13};type P struct{n string;z float64};p:=make([]P,6);q:=0.;for i:=range v{z:=math.Abs((v[i]-m[i])/s[i]);p[i]=P{names[i],z};q+=z*z};sort.Slice(p,func(i,j int)bool{return p[i].z>p[j].z});score:=math.Sqrt(q/6);re:=[]string{};for _,x:=range p{if x.z>=2{re=append(re,x.n)}};return Result{e,score,score>=3.2||p[0].z>=5,re}}
func(d *Detector)Process(e Event)Result{r:=d.Evaluate(e);d.mu.Lock();defer d.mu.Unlock();d.received++;d.processed++;if r.IsAnomaly{d.flagged=append(d.flagged,r);if len(d.flagged)>200{d.flagged=d.flagged[len(d.flagged)-200:]}};return r}
func(d *Detector)Stats()map[string]any{d.mu.RLock();defer d.mu.RUnlock();return map[string]any{"received":d.received,"processed":d.processed,"anomalies":len(d.flagged)}}
func(d *Detector)Anomalies()[]Result{d.mu.RLock();defer d.mu.RUnlock();o:=append([]Result(nil),d.flagged...);return o}

