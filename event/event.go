package event

func PutKyllingKorn(item string) string {
    // Putter kylling og korn i båt
    return "[kylling rev korn hs ---\\ \\_kylling,korn_/ _________________/---]"
}

func CrossRiverKyllingKorn(item string) string {
    // Kylling og korn krysser elva
    return "[kylling rev korn hs ---\\ _________________\\_kylling,korn_/ /---]"
}

func TakeOut(item string) string {
    // Tar ut innhold i båt
    return "[kylling rev korn hs ---\\ _________________\\__/ /---]"
}

func PutRevHs(item string) string {
    // Putter rev og hs i båt
    return "[kylling rev korn hs ---\\ \\_rev,hs_/ _________________/---]"
}

func CrossRiverRevHs(item string) string {
    // Rev krysser elva
    return "[kylling rev korn hs ---\\ _________________\\_rev,hs_/ /---]"
}