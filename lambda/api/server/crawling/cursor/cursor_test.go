package cursor

import (
	"log"
	"testing"
)

var testData = map[string][]string{
	"crawling|graph|https://facebook.com/MONOstudio-393020864107976/": {
		"https://www.facebook.com/AllDesignAgnieszkaLorenc/",
		"https://www.facebook.com/JFarchitekturawnetrz/",
		"https://www.facebook.com/KAMAODESIGN/",
		"https://www.facebook.com/MIKOLAJSKAstudio/",
		"https://www.facebook.com/NUKA-studio-1242379819212030/",
		"https://www.facebook.com/SUMA.ARCHITEKTOW/",
		"https://www.facebook.com/WonderWallStudio/",
		"https://www.facebook.com/ambienceinteriordesign/",
		"https://www.facebook.com/avocadoconcept/",
		"https://www.facebook.com/decoroom.eu/",
		"https://www.facebook.com/konzept.arch.design/",
		"https://www.facebook.com/mKosiorowskaa/",
		"https://www.facebook.com/madamemhandmade/",
		"https://www.facebook.com/pimconcept/",
		"https://www.facebook.com/projektowanie.wnetrz.krakow/",
		"https://www.facebook.com/projektowaniewnetrz.goldenpoint/",
		"https://www.facebook.com/projektowaniewnetrzmotyl/",
		"https://www.facebook.com/tillaarchitects/",
		"https://www.facebook.com/woodirongroup/",
		"https://www.facebook.com/wzstudioprojektowe/",
	},
	"crawling|graph|https://www.facebook.com/JFarchitekturawnetrz/": {
		"https://www.facebook.com/Highline.Language.School.PRO/",
		"https://www.facebook.com/JTGRUPA/",
		"https://www.facebook.com/KAZA-Concept-506844983145739/",
		"https://www.facebook.com/ambienceinteriordesign/",
		"https://www.facebook.com/architektwnetrzklaudiapniak/",
		"https://www.facebook.com/blania.art/",
		"https://www.facebook.com/elementypracownia/",
		"https://www.facebook.com/ewaprojektuje/",
		"https://www.facebook.com/formeastudio/",
		"https://www.facebook.com/hajastudio/",
		"https://www.facebook.com/kieweljanus/",
		"https://www.facebook.com/manedolnychmlynow/",
		"https://www.facebook.com/martawanat.projektowaniewnetrz/",
		"https://www.facebook.com/nie.bo.design/",
		"https://www.facebook.com/pimconcept/",
		"https://www.facebook.com/pracownia72/",
		"https://www.facebook.com/projektywstudio/",
		"https://www.facebook.com/summeragency3D/",
		"https://www.facebook.com/udzielaarchitekci/",
		"https://www.facebook.com/wzstudioprojektowe/",
	},
	"crawling|graph|https://www.facebook.com/KAMAODESIGN/": {
		"https://www.facebook.com/AFormA.interior.design/",
		"https://www.facebook.com/AllDesignAgnieszkaLorenc/",
		"https://www.facebook.com/DSwiatkowskapl/",
		"https://www.facebook.com/Gosia-Sosnowska-Photography-182070368512426/",
		"https://www.facebook.com/IV-Design-Projektowanie-wnętrz-165315113583431/",
		"https://www.facebook.com/InmagineProjektowanieWnetrz/",
		"https://www.facebook.com/JE-Projekt-Sp-z-oo-1703363736626275/",
		"https://www.facebook.com/ProjektowanieMK/",
		"https://www.facebook.com/SUMA.ARCHITEKTOW/",
		"https://www.facebook.com/StudioReneDesign/",
		"https://www.facebook.com/WarszawskieCentrumAikido/",
		"https://www.facebook.com/WszystkieWnetrzaDozwolone/",
		"https://www.facebook.com/ZawickaID/",
		"https://www.facebook.com/dzikaarchitektura/",
		"https://www.facebook.com/jasnowskacom/",
		"https://www.facebook.com/krystynaregulskaprojektant/",
		"https://www.facebook.com/livingroomstudio/",
		"https://www.facebook.com/projektowanie.wnetrz.krakow/",
		"https://www.facebook.com/projektowaniewnetrz.goldenpoint/",
		"https://www.facebook.com/rownopodsufitem/",
	},
	"crawling|graph|https://www.facebook.com/MIKOLAJSKAstudio/": {
		"https://www.facebook.com/A2studio-Pracownia-Architektury-268341369954232/",
		"https://www.facebook.com/Finchstudioarchitekturawnetrz/",
		"https://www.facebook.com/KUOOarchitects/",
		"https://www.facebook.com/KarolinaGruszeckaDietetyk/",
		"https://www.facebook.com/ProstyPlanBlog/",
		"https://www.facebook.com/WonderWallStudio/",
		"https://www.facebook.com/ambienceinteriordesign/",
		"https://www.facebook.com/architektnaszpilkach/",
		"https://www.facebook.com/bartekwlodarczykarchitekt/",
		"https://www.facebook.com/designzoo.poland/",
		"https://www.facebook.com/elementypracownia/",
		"https://www.facebook.com/foormapracownia/",
		"https://www.facebook.com/konzept.arch.design/",
		"https://www.facebook.com/kraszewska.architektura/",
		"https://www.facebook.com/manaza.interior.design/",
		"https://www.facebook.com/prostewnetrze/",
		"https://www.facebook.com/razoo-architekci-293570377331980/",
		"https://www.facebook.com/robimyzmetalu/",
		"https://www.facebook.com/studiohatch/",
		"https://www.facebook.com/wzstudioprojektowe/",
	},
	"crawling|graph|https://www.facebook.com/SUMA.ARCHITEKTOW/": {
		"https://www.facebook.com/FrantaGroup/",
		"https://www.facebook.com/KKATwnetrza/",
		"https://www.facebook.com/MONOstudio-393020864107976/",
		"https://www.facebook.com/OES-architekci-1552475751689339/",
		"https://www.facebook.com/OchAchConcept/",
		"https://www.facebook.com/ProxaYourHome/",
		"https://www.facebook.com/art-eria-189976227700742/",
		"https://www.facebook.com/bniqualitykrakow/",
		"https://www.facebook.com/concreate.design/",
		"https://www.facebook.com/ewaprojektuje/",
		"https://www.facebook.com/inwizjastudio/",
		"https://www.facebook.com/izzifurniture.eu/",
		"https://www.facebook.com/jmsSTUDIOarchitekci/",
		"https://www.facebook.com/martawypych.pracownia/",
		"https://www.facebook.com/nelajustpl/",
		"https://www.facebook.com/nieskromneprogi/",
		"https://www.facebook.com/noviartwnetrza/",
		"https://www.facebook.com/rownopodsufitem/",
		"https://www.facebook.com/tobiaszarchitekci/",
		"https://www.facebook.com/wnetrza.krakow/",
	},
	"crawling|graph|https://www.facebook.com/WonderWallStudio/": {
		"https://www.facebook.com/Archidzielo/",
		"https://www.facebook.com/CARREApl/",
		"https://www.facebook.com/GieraDesign/",
		"https://www.facebook.com/GrupaTubadzin/",
		"https://www.facebook.com/MIKOLAJSKAstudio/",
		"https://www.facebook.com/Marta-Czerkies-Home-Designer-1546380962252811/",
		"https://www.facebook.com/NowodvorskiLighting/",
		"https://www.facebook.com/PaniTapetka/",
		"https://www.facebook.com/ProstyPlanBlog/",
		"https://www.facebook.com/RawDecorPL/",
		"https://www.facebook.com/ambienceinteriordesign/",
		"https://www.facebook.com/euformasalon/",
		"https://www.facebook.com/kaflando/",
		"https://www.facebook.com/konzept.arch.design/",
		"https://www.facebook.com/miruc.krzysztof/",
		"https://www.facebook.com/mossdecor.studio/",
		"https://www.facebook.com/opainteriors/",
		"https://www.facebook.com/robimyzmetalu/",
		"https://www.facebook.com/wallartpl/",
		"https://www.facebook.com/warsawhomeexpo/",
	},
	"crawling|graph|https://www.facebook.com/ambienceinteriordesign/": {
		"https://www.facebook.com/AllDesignAgnieszkaLorenc/",
		"https://www.facebook.com/A2studio-Pracownia-Architektury-268341369954232/",
		"https://www.facebook.com/JTGRUPA/",
		"https://www.facebook.com/JUNGLEHOLIC/",
		"https://www.facebook.com/MIKOLAJSKAstudio/",
		"https://www.facebook.com/Marta-Czerkies-Home-Designer-1546380962252811/",
		"https://www.facebook.com/Nubo-837477919613833/",
		"https://www.facebook.com/Projekty-Aranżacje-Wnętrz-1029984690439965/",
		"https://www.facebook.com/ProstyPlanBlog/",
		"https://www.facebook.com/TargiHomeDecor/",
		"https://www.facebook.com/bartekwlodarczykarchitekt/",
		"https://www.facebook.com/conceptsevenstudio/",
		"https://www.facebook.com/foormapracownia/",
		"https://www.facebook.com/glamourflowerbox/",
		"https://www.facebook.com/hilightdesignofficial/",
		"https://www.facebook.com/ideabymag/",
		"https://www.facebook.com/konzept.arch.design/",
		"https://www.facebook.com/kraszewska.architektura/",
		"https://www.facebook.com/martaogrodowczykstudio/",
		"https://www.facebook.com/razoo-architekci-293570377331980/",
		"https://www.facebook.com/wzstudioprojektowe/",
	},
	"crawling|graph|https://www.facebook.com/avocadoconcept/": {
		"https://www.facebook.com/32nowastomatologia/",
		"https://www.facebook.com/A2studio-Pracownia-Architektury-268341369954232/",
		"https://www.facebook.com/DOMagalaDesign/",
		"https://www.facebook.com/Ekotektura/",
		"https://www.facebook.com/Emilkrecilody/",
		"https://www.facebook.com/Krakow.BoConcept/",
		"https://www.facebook.com/KrakowskiDziadek/",
		"https://www.facebook.com/MokaaArchitekci/",
		"https://www.facebook.com/Studio-DOMO-Dorota-i-Jarosław-Ormezowscy-462596417212840/",
		"https://www.facebook.com/byann.designstore/",
		"https://www.facebook.com/jmsSTUDIOarchitekci/",
		"https://www.facebook.com/laurazubelarchitektwnetrz/",
		"https://www.facebook.com/maxfliz/",
		"https://www.facebook.com/merapiarchitects/",
		"https://www.facebook.com/motifostudio/",
		"https://www.facebook.com/pimconcept/",
		"https://www.facebook.com/projektowaniewnetrzmotyl/",
		"https://www.facebook.com/prostewnetrze/",
		"https://www.facebook.com/superpozycja/",
		"https://www.facebook.com/wonderspacepl/",
	},
	"crawling|graph|https://www.facebook.com/pimconcept/": {
		"https://www.facebook.com/Adaventure-1562384094075209/",
		"https://www.facebook.com/BercalStudioSwiatla/",
		"https://www.facebook.com/CeramCityLUXHOME/",
		"https://www.facebook.com/JTGRUPA/",
		"https://www.facebook.com/KRES.Architekci/",
		"https://www.facebook.com/Kolektyw-317935152125045/",
		"https://www.facebook.com/NUKA-studio-1242379819212030/",
		"https://www.facebook.com/Retro-Sklep-553760011423135/",
		"https://www.facebook.com/elementypracownia/",
		"https://www.facebook.com/hajastudio/",
		"https://www.facebook.com/kijewskikoncept/",
		"https://www.facebook.com/kroniki.studio/",
		"https://www.facebook.com/polakdesign/",
		"https://www.facebook.com/pracownia72/",
		"https://www.facebook.com/projektowaniewnetrzmotyl/",
		"https://www.facebook.com/projektywstudio/",
		"https://www.facebook.com/stellarstudiopl/",
		"https://www.facebook.com/whitefoxwallpapers/",
		"https://www.facebook.com/wondershoppl/",
		"https://www.facebook.com/wzstudioprojektowe/",
	},
	"crawling|graph|https://www.facebook.com/projektowanie.wnetrz.krakow/": {
		"https://www.facebook.com/ARusecka.Architekt/",
		"https://www.facebook.com/AllDesignAgnieszkaLorenc/",
		"https://www.facebook.com/Altercoffee-Marcin-Rusnarczyk-1457373901157957/",
		"https://www.facebook.com/FreszdizajnProjektowanieWnetrz/",
		"https://www.facebook.com/Gardenbeat-Silesia-1619603568310401/",
		"https://www.facebook.com/KAMAODESIGN/",
		"https://www.facebook.com/MONOstudio-393020864107976/",
		"https://www.facebook.com/MieszkanioweMetamorfozy/",
		"https://www.facebook.com/ModneWnetrzacom/",
		"https://www.facebook.com/NiesamowiteWnetrza/",
		"https://www.facebook.com/PSJ-o-Kraków-Muzycy-Pracownicy-Przyjaciele-108422162525722/",
		"https://www.facebook.com/ProjektowanieMK/",
		"https://www.facebook.com/WszystkieWnetrzaDozwolone/",
		"https://www.facebook.com/archipuzzle/",
		"https://www.facebook.com/ejmartstudio/",
		"https://www.facebook.com/klimman.krakow/",
		"https://www.facebook.com/ogrodyaleny/",
		"https://www.facebook.com/projektowaniewnetrz.goldenpoint/",
		"https://www.facebook.com/wnetrzarki/",
		"https://www.facebook.com/zelekkarolina/",
	},
	"crawling|graph|https://www.facebook.com/projektowaniewnetrzmotyl/": {
		"https://www.facebook.com/ArchistanPracowniaArchitektoniczna/",
		"https://www.facebook.com/DOMagalaDesign/",
		"https://www.facebook.com/Gessi.Official/",
		"https://www.facebook.com/HukArchitekci/",
		"https://www.facebook.com/KKArchitekci-953841154693763/",
		"https://www.facebook.com/MIKOLAJSKAstudio/",
		"https://www.facebook.com/MONOstudio-393020864107976/",
		"https://www.facebook.com/atelierwnetrzjoannajaworska/",
		"https://www.facebook.com/euformasalon/",
		"https://www.facebook.com/formeastudio/",
		"https://www.facebook.com/garywoodpecker/",
		"https://www.facebook.com/laurazubelarchitektwnetrz/",
		"https://www.facebook.com/mKosiorowskaa/",
		"https://www.facebook.com/merapiarchitects/",
		"https://www.facebook.com/mothi.form/",
		"https://www.facebook.com/quboarchitekci/",
		"https://www.facebook.com/studiohatch/",
		"https://www.facebook.com/wnetrzaprostyuklad/",
		"https://www.facebook.com/wonderspacepl/",
		"https://www.facebook.com/zietarskaPracownia/",
	},
	"crawling|graph|https://www.facebook.com/tillaarchitects/": {
		"https://www.facebook.com/DEKAAARCHITECTS/",
		"https://www.facebook.com/InterArchArchitekci/",
		"https://www.facebook.com/KKArchitekci-953841154693763/",
		"https://www.facebook.com/KRES.Architekci/",
		"https://www.facebook.com/KUOOarchitects/",
		"https://www.facebook.com/NUKA-studio-1242379819212030/",
		"https://www.facebook.com/TK-Architekci-744842258881063/",
		"https://www.facebook.com/ambienceinteriordesign/",
		"https://www.facebook.com/bartekwlodarczykarchitekt/",
		"https://www.facebook.com/elementypracownia/",
		"https://www.facebook.com/foormapracownia/",
		"https://www.facebook.com/formeastudio/",
		"https://www.facebook.com/fussstudio/",
		"https://www.facebook.com/kraszewska.architektura/",
		"https://www.facebook.com/laurazubelarchitektwnetrz/",
		"https://www.facebook.com/madamaprojekty/",
		"https://www.facebook.com/multiwnetrza/",
		"https://www.facebook.com/razoo-architekci-293570377331980/",
		"https://www.facebook.com/studiohatch/",
		"https://www.facebook.com/wzornikarchitekta/",
	},
}

func TestCrawling(t *testing.T) {
	cursor := bfsCursor()

	if len(cursor) <= 0 {
		log.Printf("%s", "empty cursor")
		return
	}

	rootKey := "crawling|graph|" + cursor[0]
	data := testData[rootKey]
	saveGraph(cursor[0], data)

	log.Printf("%s", cursor[0])

}
