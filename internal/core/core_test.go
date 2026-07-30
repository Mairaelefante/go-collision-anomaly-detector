package core
import "testing"
func TestDetector(t *testing.T){d:=New();n:=d.Evaluate(Event{EventID:"n",TotalPT:150,MissingET:35,JetCount:2,LeptonCount:1,InvariantMass:91.2,DetectorOccupancy:.4});if n.IsAnomaly{t.Fatal("normal flagged")};a:=d.Evaluate(Event{EventID:"a",TotalPT:900,MissingET:350,JetCount:10,LeptonCount:5,InvariantMass:500,DetectorOccupancy:.98});if !a.IsAnomaly{t.Fatal("anomaly missed")}}

