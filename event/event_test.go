// Package for testing events in the riverworld
// Possible events: put(item), getin(), getout(), crossriver(), takeout(item)
// start() or setup(), reset() osv.
package event

import "testing"

//func TestPut(t *testing.T) {
//    // Hva forventer jeg?
//    wanted := "[kylling rev korn hs ---\\ \\_korn_/ _________________/---]"
//    got := Put("korn") // Hva fikk jeg?
//    if got != wanted {
//        t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
//    }
//}
func TestPutKyllingKorn(t *testing.T) {
    // Putte kylling og korn i båt
    wanted := "[kylling rev korn hs ---\\ \\_kylling,korn_/ _________________/---]"
    got := PutKyllingKorn("kylling,korn") // Hva fikk jeg?
    if got != wanted {
        t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
    }
}

func TestCrossRiverKyllingKorn(t *testing.T) {
    // Kyllin og korn krysser elva
    wanted := "[kylling rev korn hs ---\\ _________________\\_kylling,korn_/ /---]"
    got := CrossRiverKyllingKorn("kylling,korn") // Hva fikk jeg?
    if got != wanted {
        t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
    }
}

func TestTakeOut(t *testing.T) {
    // Tar ut innholdet i båt
    wanted := "[kylling rev korn hs ---\\ _________________\\__/ /---]"
    got := TakeOut("") // Hva fikk jeg?
    if got != wanted {
        t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
    }
}

func TestPutRevHs(t *testing.T) {
    // Putter rev og hs i båt
    wanted := "[kylling rev korn hs ---\\ \\_rev,hs_/ _________________/---]"
    got := PutRevHs("rev,hs") // Hva fikk jeg?
    if got != wanted {
        t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
    }
}

func TestCrossRiverRevHs(t *testing.T) {
    // Forventer at rev krysser med båt
    wanted := "[kylling rev korn hs ---\\ _________________\\_rev,hs_/ /---]"
    got := CrossRiverRevHs("rev") // Rev og hs krysser med båt
    if got != wanted {
        t.Errorf("Feil, fikk %q, ønsket %q.", got, wanted)
    }
}