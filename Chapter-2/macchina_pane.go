package main

import "fmt"

const (
	impastoPrimPaneB, impastoPrimPaneD int32 = 15, 20
	lievitazionePrimPaneB, lievitaionePrimPaneD int32 = 60, 60
	impastoSecoPaneB, impastoSecPaneD int32 = 20, 30
	lievitazioneSecPaneB, lievitazioneSecPaneD int32 = 20, 30
	modellazionePagnottaPaneB, modellazionePagnottaPaneD int32 = 2, 2
	lievitazioneFinalePaneB, lievitazioneFinalePaneD int32 = 75, 75
	cotturaPaneB, cotturaPaneD int32 = 45, 35
	raffreddamentoPaneB, raffreddamentoPaneD int32 = 30, 30
	paneBianco, paneDolce string = "B", "D"
	grandezzaNormale, grandezzaDoppia string = "N", "D"
	cotturaManuale, cotturaMacchina string = "M", "N" 
)

func main() {
	var tipologiaPane, dimensionePagnotta, tipologiaCottura string = "B", "D", "M"
	impostaPassi(tipologiaPane, dimensionePagnotta, tipologiaCottura)
}

func impostaPassi(pane string, dimPane string, tipoCottura string) {
	if pane == paneBianco {
		cuociPaneBianco(dimPane, tipoCottura)
	} else {
		cuociPaneDolce(dimPane, tipoCottura)
	}
}

func cuociPaneBianco(dimPane string, tipoCottura string) {
	if dimPane == grandezzaNormale {
		stampaProcessoNormaleB(tipoCottura)
	} else {
		stampaProcessoDoppioB(tipoCottura)
	}
}


func cuociPaneDolce(dimPane string, tipoCottura string) {
	if dimPane == grandezzaNormale {
		stampaProcessoNormaleD(tipoCottura)
	} else {
		stampaProcessoDoppioD(tipoCottura)
	}
}

func stampaProcessoNormaleB(tipoCottura string) {
	if tipoCottura == cotturaMacchina {
		tempoTotale := impastoPrimPaneB + lievitazionePrimPaneB + impastoSecoPaneB + 
			           lievitazioneSecPaneB +  modellazionePagnottaPaneB + lievitazioneFinalePaneB + 
			           cotturaPaneB + raffreddamentoPaneB
		fmt.Println("Tempo totale di cottura : %d", tempoTotale)
	} else {
		tempoTotale := impastoPrimPaneB + lievitazionePrimPaneB + impastoSecoPaneB + 
			           lievitazioneSecPaneB +  modellazionePagnottaPaneB
		fmt.Printf("Tempo Totale di Cottura : %d", tempoTotale)
	}
}

func stampaProcessoDoppioB(tipoCottura string) {
	if tipoCottura == cotturaMacchina {
		tempoCotturaRadd := (cotturaPaneB * 100) / 50
		tempoTotale := impastoPrimPaneB + lievitazionePrimPaneB + impastoSecoPaneB + 
			           lievitazioneSecPaneB +  modellazionePagnottaPaneB + lievitazioneFinalePaneB + 
			           tempoCotturaRadd + raffreddamentoPaneB
		fmt.Println("Tempo totale di cottura : %d", tempoTotale)
	} else {
		tempoTotale := impastoPrimPaneB + lievitazionePrimPaneB + impastoSecoPaneB + 
			           lievitazioneSecPaneB +  modellazionePagnottaPaneB
		fmt.Printf("Tempo Totale di Cottura : %d", tempoTotale)
	}
}

func stampaProcessoNormaleD(tipoCottura string) {
	if tipoCottura == cotturaMacchina {
		tempoTotale := impastoPrimPaneD + lievitazionePrimPaneD + impastoSecoPaneD + 
			           lievitazioneSecPaneD +  modellazionePagnottaPaneD + lievitazioneFinalePaneD + 
			           cotturaPaneD + raffreddamentoPaneD
		fmt.Printf("Tempo di cottura totale : %d", tempoTotale)
	} else {
		tempoTotale := impastoPrimPaneD + lievitazionePrimPaneD + impastoSecoPaneD + 
			           lievitazioneSecPaneD +  modellazionePagnottaPaneD
		fmt.Printf("Tempo Totale di Cottura : %d", tempoTotale)
	}
}

func stampaProcessoDoppioD(tipoCottura string) {
	if tipoCottura == cotturaMacchina {
		tempoCotturaRadd := (cotturaPaneD * 100) / 50
		tempoTotale := impastoPrimPaneD + lievitazionePrimPaneD + impastoSecoPaneD + 
			           lievitazioneSecPaneD +  modellazionePagnottaPaneD + lievitazioneFinalePaneD + 
			           tempoCotturaRadd + raffreddamentoPaneD
		fmt.Printf("Tempo di cottura totale : %d", tempoTotale)
	} else {
		tempoTotale := impastoPrimPaneD + lievitazionePrimPaneD + impastoSecoPaneD + 
			           lievitazioneSecPaneD +  modellazionePagnottaPaneD
		fmt.Printf("Tempo Totale di Cottura : %d", tempoTotale)
	}
}