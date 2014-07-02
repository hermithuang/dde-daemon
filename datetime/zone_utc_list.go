package datetime

import "path"
import dutils "pkg.linuxdeepin.com/lib/utils"
import . "pkg.linuxdeepin.com/lib/gettext"

type zoneCityInfo struct {
	Zone string
	Desc string
}

var zoneInfos []zoneCityInfo

func checkZoneValid(zone, desc string) {
	if !dutils.IsFileExist(path.Join("/usr/share/zoneinfo", zone)) {
		return
	}

	info := zoneCityInfo{zone, desc}
	zoneInfos = append(zoneInfos, info)
}

func initZoneInfos() {
	checkZoneValid("Pacific/Marquesas", "UTC-11 "+Tr("Marquesas"))
	checkZoneValid("Pacific/Niue", "UTC-11 "+Tr("Niue"))
	checkZoneValid("US/Samoa", "UTC-11 "+Tr("Samoa"))
	checkZoneValid("America/Adak", "UTC-10 "+Tr("Adak"))
	checkZoneValid("America/Atka", "UTC-10 "+Tr("Atka"))
	checkZoneValid("HST", "UTC-10 "+Tr("HST"))
	checkZoneValid("Pacific/Honolulu", "UTC-10 "+Tr("Honolulu"))
	checkZoneValid("Pacific/Johnston", "UTC-10 "+Tr("Johnston"))
	checkZoneValid("Pacific/Rarotonga", "UTC-10 "+Tr("Rarotonga"))
	checkZoneValid("Pacific/Tahiti", "UTC-10 "+Tr("Tahiti"))
	checkZoneValid("US/Aleutian", "UTC-10 "+Tr("Aleutian"))
	checkZoneValid("US/Hawaii", "UTC-10 "+Tr("Hawaii"))
	checkZoneValid("America/Anchorage", "UTC-09 "+Tr("Anchorage"))
	checkZoneValid("America/Juneau", "UTC-09 "+Tr("Juneau"))
	checkZoneValid("America/Nome", "UTC-09 "+Tr("Nome"))
	checkZoneValid("America/Sitka", "UTC-09 "+Tr("Sitka"))
	checkZoneValid("America/Yakutat", "UTC-09 "+Tr("Yakutat"))
	checkZoneValid("Pacific/Gambier", "UTC-09 "+Tr("Gambier"))
	checkZoneValid("US/Alaska", "UTC-09 "+Tr("Alaska"))
	checkZoneValid("America/Dawson", "UTC-08 "+Tr("Dawson"))
	checkZoneValid("America/Ensenada", "UTC-08 "+Tr("Ensenada"))
	checkZoneValid("America/Los_Angeles", "UTC-08 "+Tr("Los Angeles"))
	checkZoneValid("America/Metlakatla", "UTC-08 "+Tr("Metlakatla"))
	checkZoneValid("America/Santa_Isabel", "UTC-08 "+Tr("Santa Isabel"))
	checkZoneValid("America/Tijuana", "UTC-08 "+Tr("Tijuana"))
	checkZoneValid("America/Vancouver", "UTC-08 "+Tr("Vancouver"))
	checkZoneValid("America/Whitehorse", "UTC-08 "+Tr("Whitehorse"))
	checkZoneValid("Canada/Newfoundland", "UTC-08 "+Tr("Newfoundland"))
	checkZoneValid("Canada/Yukon", "UTC-08 "+Tr("Yukon"))
	checkZoneValid("Mexico/BajaNorte", "UTC-08 "+Tr("BajaNorte"))
	checkZoneValid("Pacific/Pitcairn", "UTC-08 "+Tr("Pitcairn"))
	checkZoneValid("US/Pacific", "UTC-08 "+Tr("Pacific"))
	checkZoneValid("US/Pacific-New", "UTC-08 "+Tr("Pacific New"))
	checkZoneValid("America/Boise", "UTC-07 "+Tr("Boise"))
	checkZoneValid("America/Cambridge_Bay", "UTC-07 "+Tr("Cambridge Bay"))
	checkZoneValid("America/Chihuahua", "UTC-07 "+Tr("Chihuahua"))
	checkZoneValid("America/Creston", "UTC-07 "+Tr("Creston"))
	checkZoneValid("America/Dawson_Creek", "UTC-07 "+Tr("Dawson Creek"))
	checkZoneValid("America/Denver", "UTC-07 "+Tr("Denver"))
	checkZoneValid("America/Edmonton", "UTC-07 "+Tr("Edmonton"))
	checkZoneValid("America/Hermosillo", "UTC-07 "+Tr("Hermosillo"))
	checkZoneValid("America/Inuvik", "UTC-07 "+Tr("Inuvik"))
	checkZoneValid("America/Mazatlan", "UTC-07 "+Tr("Mazatlan"))
	checkZoneValid("America/Ojinaga", "UTC-07 "+Tr("Ojinaga"))
	checkZoneValid("America/Phoenix", "UTC-07 "+Tr("Phoenix"))
	checkZoneValid("America/Shiprock", "UTC-07 "+Tr("Shiprock"))
	checkZoneValid("America/Yellowknife", "UTC-07 "+Tr("Yellowknife"))
	checkZoneValid("Canada/Mountain", "UTC-07 "+Tr("Canada Mountain"))
	checkZoneValid("Mexico/BajaSur", "UTC-07 "+Tr("BajaSur"))
	checkZoneValid("MST", "UTC-07 "+Tr("MST"))
	checkZoneValid("Navajo", "UTC-07 "+Tr("Navajo"))
	checkZoneValid("US/Arizona", "UTC-07 "+Tr("Arizona"))
	checkZoneValid("US/Mountain", "UTC-07 "+Tr("USA Mountain"))
	checkZoneValid("America/Bahia_Banderas", "UTC-06 "+Tr("Bahia Banderas"))
	checkZoneValid("America/Belize", "UTC-06 "+Tr("Belize"))
	checkZoneValid("America/Cancun", "UTC-06 "+Tr("Cancun"))
	checkZoneValid("America/Chicago", "UTC-06 "+Tr("Chicago"))
	checkZoneValid("America/Costa_Rica", "UTC-06 "+Tr("Costa Rica"))
	checkZoneValid("America/El_Salvador", "UTC-06 "+Tr("El Salvador"))
	checkZoneValid("America/Guatemala", "UTC-06 "+Tr("Guatemala"))
	checkZoneValid("America/Indiana/Knox", "UTC-06 "+Tr("Knox"))
	checkZoneValid("America/Indiana/Tell_City", "UTC-06 "+Tr("Tell City"))
	checkZoneValid("America/Managua", "UTC-06 "+Tr("Managua"))
	checkZoneValid("America/Matamoros", "UTC-06 "+Tr("Matamoros"))
	checkZoneValid("America/Menominee", "UTC-06 "+Tr("Menominee"))
	checkZoneValid("America/Merida", "UTC-06 "+Tr("Merida"))
	checkZoneValid("America/Mexico_City", "UTC-06 "+Tr("Mexico City"))
	checkZoneValid("America/Monterrey", "UTC-06 "+Tr("Monterrey"))
	checkZoneValid("America/North_Dakota/Beulah", "UTC-06 "+Tr("Beulah"))
	checkZoneValid("America/North_Dakota/New_Salem", "UTC-06 "+Tr("New Salem"))
	checkZoneValid("America/Rainy_River", "UTC-06 "+Tr("Rainy River"))
	checkZoneValid("America/Swift_Current", "UTC-06 "+Tr("Swift Current"))
	checkZoneValid("America/Regina", "UTC-06 "+Tr("Regina"))
	checkZoneValid("America/Resolute", "UTC-06 "+Tr("Resolute"))
	checkZoneValid("America/Tegucigalpa", "UTC-06 "+Tr("Tegucigalpa"))
	checkZoneValid("America/Winnipeg", "UTC-06 "+Tr("Winnipeg"))
	checkZoneValid("Canada/East-Saskatchewan", "UTC-06 "+Tr("East Saskatchewan"))
	checkZoneValid("Canada/Saskatchewan", "UTC-06 "+Tr("Saskatchewan"))
	checkZoneValid("Chile/EasterIsland", "UTC-06 "+Tr("EasterIsland"))
	checkZoneValid("Mexico/General", "UTC-06 "+Tr("General"))
	checkZoneValid("Pacific/Easter", "UTC-06 "+Tr("Pacific East"))
	checkZoneValid("Pacific/Galapagos", "UTC-06 "+Tr("Galapagos"))
	checkZoneValid("US/Central", "UTC-06 "+Tr("USA Central"))
	checkZoneValid("US/Indiana-Starke", "UTC-06 "+Tr("Indiana Starke"))
	checkZoneValid("America/Atikokan", "UTC-05 "+Tr("Atikokan"))
	checkZoneValid("America/Bogota", "UTC-05 "+Tr("Bogota"))
	checkZoneValid("America/Cayman", "UTC-05 "+Tr("Cayman"))
	checkZoneValid("America/Coral_Harbour", "UTC-05 "+Tr("Coral Harbour"))
	checkZoneValid("America/Detroit", "UTC-05 "+Tr("Detroit"))
	checkZoneValid("America/Fort_Wayne", "UTC-05 "+Tr("Fort Wayne"))
	checkZoneValid("America/Grand_Turk", "UTC-05 "+Tr("Grand Turk"))
	checkZoneValid("America/Guayaquil", "UTC-05 "+Tr("Guayaquil"))
	checkZoneValid("America/Havana", "UTC-05 "+Tr("Havana"))
	checkZoneValid("America/Indiana/Marengo", "UTC-05 "+Tr("Marengo"))
	checkZoneValid("America/Indiana/Petersburg", "UTC-05 "+Tr("Petersburg"))
	checkZoneValid("America/Indiana/Vevay", "UTC-05 "+Tr("Vevay"))
	checkZoneValid("America/Indiana/Vincennes", "UTC-05 "+Tr("Vincennes"))
	checkZoneValid("America/Indiana/Winamac", "UTC-05 "+Tr("Winamac"))
	checkZoneValid("America/Indianapolis", "UTC-05 "+Tr("Indianapolis"))
	checkZoneValid("America/Iqaluit", "UTC-05 "+Tr("Iqaluit"))
	checkZoneValid("America/Jamaica", "UTC-05 "+Tr("Jamaica"))
	checkZoneValid("America/Lima", "UTC-05 "+Tr("Lima"))
	checkZoneValid("America/Louisville", "UTC-05 "+Tr("Louisville"))
	checkZoneValid("America/Kentucky/Monticello", "UTC-05 "+Tr("Monticello"))
	checkZoneValid("America/Montreal", "UTC-05 "+Tr("Montreal"))
	checkZoneValid("America/Nassau", "UTC-05 "+Tr("Nassau"))
	checkZoneValid("America/New_York", "UTC-05 "+Tr("New York"))
	checkZoneValid("America/Nipigon", "UTC-05 "+Tr("Nipigon"))
	checkZoneValid("America/Panama", "UTC-05 "+Tr("Panama"))
	checkZoneValid("America/Port-au-Prince", "UTC-05 "+Tr("Port au Prince"))
	checkZoneValid("America/Porto_Acre", "UTC-05 "+Tr("Porto Acre"))
	checkZoneValid("America/Rio_Branco", "UTC-05 "+Tr("Rio Branco"))
	checkZoneValid("America/Thunder_Bay", "UTC-05 "+Tr("Thunder Bay"))
	checkZoneValid("America/Toronto", "UTC-05 "+Tr("Toronto"))
	checkZoneValid("Australia/Yancowinna", "UTC-05 "+Tr("Yancowinna"))
	checkZoneValid("Canada/Eastern", "UTC-05 "+Tr("Canada Eastern"))
	checkZoneValid("Cuba", "UTC-05 "+Tr("Cuba"))
	checkZoneValid("EST", "UTC-05 "+Tr("EST"))
	checkZoneValid("US/Eastern", "UTC-05 "+Tr("USA Eastern"))
	checkZoneValid("US/East-Indiana", "UTC-05 "+Tr("East Indiana"))
	checkZoneValid("US/Michigan", "UTC-05 "+Tr("Michigan"))
	checkZoneValid("America/Anguilla", "UTC-04 "+Tr("Anguilla"))
	checkZoneValid("America/Antigua", "UTC-04 "+Tr("Antigua"))
	checkZoneValid("America/Aruba", "UTC-04 "+Tr("Aruba"))
	checkZoneValid("America/Asuncion", "UTC-04 "+Tr("Asuncion"))
	checkZoneValid("America/Barbados", "UTC-04 "+Tr("Barbados"))
	checkZoneValid("America/Boa_Vista", "UTC-04 "+Tr("Boa Vista"))
	checkZoneValid("America/Campo_Grande", "UTC-04 "+Tr("Campo Grande"))
	checkZoneValid("America/Cuiaba", "UTC-04 "+Tr("Cuiaba"))
	checkZoneValid("America/Curacao", "UTC-04 "+Tr("Curacao"))
	checkZoneValid("America/Dominica", "UTC-04 "+Tr("Dominica"))
	checkZoneValid("America/Eirunepe", "UTC-04 "+Tr("Eirunepe"))
	checkZoneValid("America/Glace_Bay", "UTC-04 "+Tr("Glace Bay"))
	checkZoneValid("America/Goose_Bay", "UTC-04 "+Tr("Goose Bay"))
	checkZoneValid("America/Grenada", "UTC-04 "+Tr("Grenada"))
	checkZoneValid("America/Guadeloupe", "UTC-04 "+Tr("Guadeloupe"))
	checkZoneValid("America/Guyana", "UTC-04 "+Tr("Guyana"))
	checkZoneValid("America/Halifax", "UTC-04 "+Tr("Halifax"))
	checkZoneValid("America/Kralendijk", "UTC-04 "+Tr("Kralendijk"))
	checkZoneValid("America/La_Paz", "UTC-04 "+Tr("La Paz"))
	checkZoneValid("America/Lower_Princes", "UTC-04 "+Tr("Lower Princes"))
	checkZoneValid("America/Manaus", "UTC-04 "+Tr("Manaus"))
	checkZoneValid("America/Marigot", "UTC-04 "+Tr("Marigot"))
	checkZoneValid("America/Martinique", "UTC-04 "+Tr("Martinique"))
	checkZoneValid("America/Moncton", "UTC-04 "+Tr("Moncton"))
	checkZoneValid("America/Montserrat", "UTC-04 "+Tr("Montserrat"))
	checkZoneValid("America/Port_of_Spain", "UTC-04 "+Tr("Port of Spain"))
	checkZoneValid("America/Porto_Velho", "UTC-04 "+Tr("Porto Velho"))
	checkZoneValid("America/Puerto_Rico", "UTC-04 "+Tr("Puerto Rico"))
	checkZoneValid("America/Santiago", "UTC-04 "+Tr("Santiago"))
	checkZoneValid("America/Santo_Domingo", "UTC-04 "+Tr("Santo Domingo"))
	checkZoneValid("America/St_Barthelemy", "UTC-04 "+Tr("St Barthelemy"))
	checkZoneValid("America/St_Johns", "UTC-04 "+Tr("St Johns"))
	checkZoneValid("America/St_Lucia", "UTC-04 "+Tr("St Lucia"))
	checkZoneValid("America/St_Thomas", "UTC-04 "+Tr("St Thomas"))
	checkZoneValid("America/St_Vincent", "UTC-04 "+Tr("St Vincent"))
	checkZoneValid("America/Thule", "UTC-04 "+Tr("Thule"))
	checkZoneValid("America/Tortola", "UTC-04 "+Tr("Tortola"))
	checkZoneValid("America/Virgin", "UTC-04 "+Tr("Virgin"))
	checkZoneValid("Antarctica/Palmer", "UTC-04 "+Tr("Palmer"))
	checkZoneValid("Atlantic/Bermuda", "UTC-04 "+Tr("Bermuda"))
	checkZoneValid("Brazil/West", "UTC-04 "+Tr("Brazil West"))
	checkZoneValid("Canada/Atlantic", "UTC-04 "+Tr("Atlantic"))
	checkZoneValid("Chile/Continental", "UTC-04 "+Tr("Continental"))
	checkZoneValid("America/Araguaina", "UTC-03 "+Tr("Araguaina"))
	checkZoneValid("America/Argentina/Catamarca", "UTC-03 "+Tr("Catamarca"))
	checkZoneValid("America/Argentina/ComodRivadavia", "UTC-03 "+Tr("ComodRivadavia"))
	checkZoneValid("America/Argentina/Jujuy", "UTC-03 "+Tr("Jujuy"))
	checkZoneValid("America/Argentina/La_Rioja", "UTC-03 "+Tr("La Rioja"))
	checkZoneValid("America/Argentina/Rio_Gallegos", "UTC-03 "+Tr("Rio Gallegos"))
	checkZoneValid("America/Argentina/Salta", "UTC-03 "+Tr("Salta"))
	checkZoneValid("America/Argentina/San_Juan", "UTC-03 "+Tr("San Juan"))
	checkZoneValid("America/Argentina/San_Luis", "UTC-03 "+Tr("San Luis"))
	checkZoneValid("America/Argentina/Tucuman", "UTC-03 "+Tr("Tucuman"))
	checkZoneValid("America/Argentina/Ushuaia", "UTC-03 "+Tr("Ushuaia"))
	checkZoneValid("America/Bahia", "UTC-03 "+Tr("Bahia"))
	checkZoneValid("America/Belem", "UTC-03 "+Tr("Belem"))
	checkZoneValid("America/Caracas", "UTC-03 "+Tr("Caracas"))
	checkZoneValid("America/Cayenne", "UTC-03 "+Tr("Cayenne"))
	checkZoneValid("America/Cordoba", "UTC-03 "+Tr("Cordoba"))
	checkZoneValid("America/Fortaleza", "UTC-03 "+Tr("Fortaleza"))
	checkZoneValid("America/Godthab", "UTC-03 "+Tr("Godthab"))
	checkZoneValid("America/Maceio", "UTC-03 "+Tr("Maceio"))
	checkZoneValid("America/Mendoza", "UTC-03 "+Tr("Mendoza"))
	checkZoneValid("America/Miquelon", "UTC-03 "+Tr("Miquelon"))
	checkZoneValid("America/Montevideo", "UTC-03 "+Tr("Montevideo"))
	checkZoneValid("America/Paramaribo", "UTC-03 "+Tr("Paramaribo"))
	checkZoneValid("America/Recife", "UTC-03 "+Tr("Recife"))
	checkZoneValid("America/Rosario", "UTC-03 "+Tr("Rosario"))
	checkZoneValid("America/Santarem", "UTC-03 "+Tr("Santarem"))
	checkZoneValid("America/Sao_Paulo", "UTC-03 "+Tr("Sao Paulo"))
	checkZoneValid("Antarctica/Rothera", "UTC-03 "+Tr("Rothera"))
	checkZoneValid("Atlantic/Stanley", "UTC-03 "+Tr("Stanley"))
	checkZoneValid("Brazil/East", "UTC-03 "+Tr("Brazil East"))
	checkZoneValid("America/Noronha", "UTC-02 "+Tr("Noronha"))
	checkZoneValid("Atlantic/South_Georgia", "UTC-02 "+Tr("South Georgia"))
	checkZoneValid("Atlantic/Cape_Verde", "UTC-01 "+Tr("Cape Verde"))
	checkZoneValid("Africa/Abidjan", "UTC+00 "+Tr("Abidjan"))
	checkZoneValid("Africa/Accra", "UTC+00 "+Tr("Accra"))
	checkZoneValid("Africa/Bamako", "UTC+00 "+Tr("Bamako"))
	checkZoneValid("Africa/Banjul", "UTC+00 "+Tr("Banjul"))
	checkZoneValid("Africa/Bissau", "UTC+00 "+Tr("Bissau"))
	checkZoneValid("Africa/Casablanca", "UTC+00 "+Tr("Casablanca"))
	checkZoneValid("Africa/Conakry", "UTC+00 "+Tr("Conakry"))
	checkZoneValid("Africa/Dakar", "UTC+00 "+Tr("Dakar"))
	checkZoneValid("Africa/El_Aaiun", "UTC+00 "+Tr("El Aaiun"))
	checkZoneValid("Africa/Freetown", "UTC+00 "+Tr("Freetown"))
	checkZoneValid("Africa/Lome", "UTC+00 "+Tr("Lome"))
	checkZoneValid("Africa/Monrovia", "UTC+00 "+Tr("Monrovia"))
	checkZoneValid("Africa/Nouakchott", "UTC+00 "+Tr("Nouakchott"))
	checkZoneValid("Africa/Ouagadougou", "UTC+00 "+Tr("Ouagadougou"))
	checkZoneValid("Africa/Sao_Tome", "UTC+00 "+Tr("Sao Tome"))
	checkZoneValid("Africa/Timbuktu", "UTC+00 "+Tr("Timbuktu"))
	checkZoneValid("Atlantic/Canary", "UTC+00 "+Tr("Canary"))
	checkZoneValid("Atlantic/Madeira", "UTC+00 "+Tr("Madeira"))
	checkZoneValid("Atlantic/Reykjavik", "UTC+00 "+Tr("Reykjavik"))
	checkZoneValid("Atlantic/St_Helena", "UTC+00 "+Tr("St Helena"))
	checkZoneValid("Eire", "UTC+00 "+Tr("Eire"))
	checkZoneValid("Europe/Belfast", "UTC+00 "+Tr("Belfast"))
	checkZoneValid("Europe/Dublin", "UTC+00 "+Tr("Dublin"))
	checkZoneValid("Europe/Guernsey", "UTC+00 "+Tr("Guernsey"))
	checkZoneValid("Europe/Isle_of_Man", "UTC+00 "+Tr("Isle of Man"))
	checkZoneValid("Europe/Jersey", "UTC+00 "+Tr("Jersey"))
	checkZoneValid("Europe/Lisbon", "UTC+00 "+Tr("Lisbon"))
	checkZoneValid("Europe/London", "UTC+00 "+Tr("London"))
	checkZoneValid("GB", "UTC+00 "+Tr("GB"))
	checkZoneValid("GB-Eire", "UTC+00 "+Tr("GB Eire"))
	checkZoneValid("GMT", "UTC+00 "+Tr("USA GMT"))
	checkZoneValid("GMT+0", "UTC+00 "+Tr("Brazil GMT"))
	checkZoneValid("Greenwich", "UTC+00 "+Tr("Greenwich"))
	checkZoneValid("Iceland", "UTC+00 "+Tr("Iceland"))
	checkZoneValid("Portugal", "UTC+00 "+Tr("Portugal"))
	checkZoneValid("UCT", "UTC+00 "+Tr("UCT"))
	checkZoneValid("Universal", "UTC+00 "+Tr("Universal"))
	checkZoneValid("UTC", "UTC+00 "+Tr("UTC"))
	checkZoneValid("WET", "UTC+00 "+Tr("WET"))
	checkZoneValid("Zulu", "UTC+00 "+Tr("Zulu"))
	checkZoneValid("Africa/Algiers", "UTC+01 "+Tr("Algiers"))
	checkZoneValid("Africa/Bangui", "UTC+01 "+Tr("Bangui"))
	checkZoneValid("Africa/Brazzaville", "UTC+01 "+Tr("Brazzaville"))
	checkZoneValid("Africa/Ceuta", "UTC+01 "+Tr("Ceuta"))
	checkZoneValid("Africa/Douala", "UTC+01 "+Tr("Douala"))
	checkZoneValid("Africa/Kinshasa", "UTC+01 "+Tr("Kinshasa"))
	checkZoneValid("Africa/Lagos", "UTC+01 "+Tr("Lagos"))
	checkZoneValid("Africa/Libreville", "UTC+01 "+Tr("Libreville"))
	checkZoneValid("Africa/Luanda", "UTC+01 "+Tr("Luanda"))
	checkZoneValid("Africa/Malabo", "UTC+01 "+Tr("Malabo"))
	checkZoneValid("Africa/Ndjamena", "UTC+01 "+Tr("Ndjamena"))
	checkZoneValid("Africa/Niamey", "UTC+01 "+Tr("Niamey"))
	checkZoneValid("Africa/Porto-Novo", "UTC+01 "+Tr("Porto Novo"))
	checkZoneValid("Africa/Tunis", "UTC+01 "+Tr("Tunis"))
	checkZoneValid("Africa/Windhoek", "UTC+01 "+Tr("Windhoek"))
	checkZoneValid("Arctic/Longyearbyen", "UTC+01 "+Tr("Longyearbyen"))
	checkZoneValid("CET", "UTC+01 "+Tr("CET"))
	checkZoneValid("Europe/Amsterdam", "UTC+01 "+Tr("Amsterdam"))
	checkZoneValid("Europe/Andorra", "UTC+01 "+Tr("Andorra"))
	checkZoneValid("Europe/Belgrade", "UTC+01 "+Tr("Belgrade"))
	checkZoneValid("Europe/Berlin", "UTC+01 "+Tr("Berlin"))
	checkZoneValid("Europe/Bratislava", "UTC+01 "+Tr("Bratislava"))
	checkZoneValid("Europe/Brussels", "UTC+01 "+Tr("Brussels"))
	checkZoneValid("Europe/Budapest", "UTC+01 "+Tr("Budapest"))
	checkZoneValid("Europe/Copenhagen", "UTC+01 "+Tr("Copenhagen"))
	checkZoneValid("Europe/Gibraltar", "UTC+01 "+Tr("Gibraltar"))
	checkZoneValid("Europe/Ljubljana", "UTC+01 "+Tr("Ljubljana"))
	checkZoneValid("Europe/Luxembourg", "UTC+01 "+Tr("Luxembourg"))
	checkZoneValid("Europe/Madrid", "UTC+01 "+Tr("Madrid"))
	checkZoneValid("Europe/Malta", "UTC+01 "+Tr("Malta"))
	checkZoneValid("Europe/Monaco", "UTC+01 "+Tr("Monaco"))
	checkZoneValid("Europe/Oslo", "UTC+01 "+Tr("Oslo"))
	checkZoneValid("Europe/Paris", "UTC+01 "+Tr("Paris"))
	checkZoneValid("Europe/Podgorica", "UTC+01 "+Tr("Podgorica"))
	checkZoneValid("Europe/Prague", "UTC+01 "+Tr("Prague"))
	checkZoneValid("Europe/Rome", "UTC+01 "+Tr("Rome"))
	checkZoneValid("Europe/San_Marino", "UTC+01 "+Tr("San Marino"))
	checkZoneValid("Europe/Sarajevo", "UTC+01 "+Tr("Sarajevo"))
	checkZoneValid("Europe/Skopje", "UTC+01 "+Tr("Skopje"))
	checkZoneValid("Europe/Stockholm", "UTC+01 "+Tr("Stockholm"))
	checkZoneValid("Europe/Tirane", "UTC+01 "+Tr("Tirane"))
	checkZoneValid("Europe/Vaduz", "UTC+01 "+Tr("Vaduz"))
	checkZoneValid("Europe/Vatican", "UTC+01 "+Tr("Vatican"))
	checkZoneValid("Europe/Vienna", "UTC+01 "+Tr("Vienna"))
	checkZoneValid("Europe/Warsaw", "UTC+01 "+Tr("Warsaw"))
	checkZoneValid("Europe/Zagreb", "UTC+01 "+Tr("Zagreb"))
	checkZoneValid("Europe/Zurich", "UTC+01 "+Tr("Zurich"))
	checkZoneValid("MET", "UTC+01 "+Tr("MET"))
	checkZoneValid("Africa/Blantyre", "UTC+02 "+Tr("Blantyre"))
	checkZoneValid("Africa/Bujumbura", "UTC+02 "+Tr("Bujumbura"))
	checkZoneValid("Africa/Cairo", "UTC+02 "+Tr("Cairo"))
	checkZoneValid("Africa/Gaborone", "UTC+02 "+Tr("Gaborone"))
	checkZoneValid("Africa/Harare", "UTC+02 "+Tr("Harare"))
	checkZoneValid("Africa/Johannesburg", "UTC+02 "+Tr("Johannesburg"))
	checkZoneValid("Africa/Kigali", "UTC+02 "+Tr("Kigali"))
	checkZoneValid("Africa/Lubumbashi", "UTC+02 "+Tr("Lubumbashi"))
	checkZoneValid("Africa/Lusaka", "UTC+02 "+Tr("Lusaka"))
	checkZoneValid("Africa/Maputo", "UTC+02 "+Tr("Maputo"))
	checkZoneValid("Africa/Maseru", "UTC+02 "+Tr("Maseru"))
	checkZoneValid("Africa/Mbabane", "UTC+02 "+Tr("Mbabane"))
	checkZoneValid("Africa/Tripoli", "UTC+02 "+Tr("Tripoli"))
	checkZoneValid("Asia/Amman", "UTC+02 "+Tr("Amman"))
	checkZoneValid("Asia/Beirut", "UTC+02 "+Tr("Beirut"))
	checkZoneValid("Asia/Damascus", "UTC+02 "+Tr("Damascus"))
	checkZoneValid("Asia/Gaza", "UTC+02 "+Tr("Gaza"))
	checkZoneValid("Asia/Hebron", "UTC+02 "+Tr("Hebron"))
	checkZoneValid("Asia/Jerusalem", "UTC+02 "+Tr("Jerusalem"))
	checkZoneValid("Asia/Tehran", "UTC+02 "+Tr("Tehran"))
	checkZoneValid("EET", "UTC+02 "+Tr("EET"))
	checkZoneValid("Egypt", "UTC+02 "+Tr("Egypt"))
	checkZoneValid("Europe/Athens", "UTC+02 "+Tr("Athens"))
	checkZoneValid("Europe/Bucharest", "UTC+02 "+Tr("Bucharest"))
	checkZoneValid("Europe/Chisinau", "UTC+02 "+Tr("Chisinau"))
	checkZoneValid("Europe/Helsinki", "UTC+02 "+Tr("Helsinki"))
	checkZoneValid("Europe/Istanbul", "UTC+02 "+Tr("Istanbul"))
	checkZoneValid("Europe/Kiev", "UTC+02 "+Tr("Kiev"))
	checkZoneValid("Europe/Mariehamn", "UTC+02 "+Tr("Mariehamn"))
	checkZoneValid("Europe/Nicosia", "UTC+02 "+Tr("Nicosia"))
	checkZoneValid("Europe/Riga", "UTC+02 "+Tr("Riga"))
	checkZoneValid("Europe/Simferopol", "UTC+02 "+Tr("Simferopol"))
	checkZoneValid("Europe/Sofia", "UTC+02 "+Tr("Sofia"))
	checkZoneValid("Europe/Tallinn", "UTC+02 "+Tr("Tallinn"))
	checkZoneValid("Europe/Uzhgorod", "UTC+02 "+Tr("Uzhgorod"))
	checkZoneValid("Europe/Vilnius", "UTC+02 "+Tr("Vilnius"))
	checkZoneValid("Europe/Zaporozhye", "UTC+02 "+Tr("Zaporozhye"))
	checkZoneValid("Iran", "UTC+02 "+Tr("Iran"))
	checkZoneValid("Libya", "UTC+02 "+Tr("Libya"))
	checkZoneValid("Turkey", "UTC+02 "+Tr("Turkey"))
	checkZoneValid("Africa/Addis_Ababa", "UTC+03 "+Tr("Addis Ababa"))
	checkZoneValid("Africa/Asmara", "UTC+03 "+Tr("Asmara"))
	checkZoneValid("Africa/Dar_es_Salaam", "UTC+03 "+Tr("Dar es Salaam"))
	checkZoneValid("Africa/Djibouti", "UTC+03 "+Tr("Djibouti"))
	checkZoneValid("Africa/Juba", "UTC+03 "+Tr("Juba"))
	checkZoneValid("Africa/Kampala", "UTC+03 "+Tr("Kampala"))
	checkZoneValid("Africa/Khartoum", "UTC+03 "+Tr("Khartoum"))
	checkZoneValid("Africa/Mogadishu", "UTC+03 "+Tr("Mogadishu"))
	checkZoneValid("Africa/Nairobi", "UTC+03 "+Tr("Nairobi"))
	checkZoneValid("Asia/Aden", "UTC+03 "+Tr("Aden"))
	checkZoneValid("Asia/Baghdad", "UTC+03 "+Tr("Baghdad"))
	checkZoneValid("Asia/Bahrain", "UTC+03 "+Tr("Bahrain"))
	checkZoneValid("Asia/Kuwait", "UTC+03 "+Tr("Kuwait"))
	checkZoneValid("Asia/Rangoon", "UTC+03 "+Tr("Rangoon"))
	checkZoneValid("Asia/Qatar", "UTC+03 "+Tr("Qatar"))
	checkZoneValid("Europe/Kaliningrad", "UTC+03 "+Tr("Kaliningrad"))
	checkZoneValid("Europe/Minsk", "UTC+03 "+Tr("Minsk"))
	checkZoneValid("Indian/Antananarivo", "UTC+03 "+Tr("Antananarivo"))
	checkZoneValid("Asia/Baku", "UTC+04 "+Tr("Baku"))
	checkZoneValid("Asia/Dubai", "UTC+04 "+Tr("Dubai"))
	checkZoneValid("Asia/Muscat", "UTC+04 "+Tr("Muscat"))
	checkZoneValid("Asia/Tbilisi", "UTC+04 "+Tr("Tbilisi"))
	checkZoneValid("Asia/Yerevan", "UTC+04 "+Tr("Yerevan"))
	checkZoneValid("Europe/Moscow", "UTC+04 "+Tr("Moscow"))
	checkZoneValid("Europe/Samara", "UTC+04 "+Tr("Samara"))
	checkZoneValid("Europe/Volgograd", "UTC+04 "+Tr("Volgograd"))
	checkZoneValid("Indian/Mahe", "UTC+04 "+Tr("Mahe"))
	checkZoneValid("Indian/Mauritius", "UTC+04 "+Tr("Mauritius"))
	checkZoneValid("Indian/Reunion", "UTC+04 "+Tr("Reunion"))
	checkZoneValid("Antarctica/Davis", "UTC+05 "+Tr("Davis"))
	checkZoneValid("Antarctica/Mawson", "UTC+05 "+Tr("Mawson"))
	checkZoneValid("Asia/Aqtau", "UTC+05 "+Tr("Aqtau"))
	checkZoneValid("Asia/Aqtobe", "UTC+05 "+Tr("Aqtobe"))
	checkZoneValid("Asia/Ashkhabad", "UTC+05 "+Tr("Ashkhabad"))
	checkZoneValid("Asia/Dushanbe", "UTC+05 "+Tr("Dushanbe"))
	checkZoneValid("Asia/Karachi", "UTC+05 "+Tr("Karachi"))
	checkZoneValid("Asia/Oral", "UTC+05 "+Tr("Oral"))
	checkZoneValid("Asia/Samarkand", "UTC+05 "+Tr("Samarkand"))
	checkZoneValid("Asia/Tashkent", "UTC+05 "+Tr("Tashkent"))
	checkZoneValid("Indian/Kerguelen", "UTC+05 "+Tr("Kerguelen"))
	checkZoneValid("Indian/Maldives", "UTC+05 "+Tr("Maldives"))
	checkZoneValid("Antarctica/Vostok", "UTC+06 "+Tr("Vostok"))
	checkZoneValid("Asia/Almaty", "UTC+06 "+Tr("Almaty"))
	checkZoneValid("Asia/Bishkek", "UTC+06 "+Tr("Bishkek"))
	checkZoneValid("Asia/Colombo", "UTC+06 "+Tr("Colombo"))
	checkZoneValid("Asia/Dhaka", "UTC+06 "+Tr("Dhaka"))
	checkZoneValid("Asia/Qyzylorda", "UTC+06 "+Tr("Qyzylorda"))
	checkZoneValid("Asia/Thimphu", "UTC+06 "+Tr("Thimphu"))
	checkZoneValid("Asia/Yekaterinburg", "UTC+06 "+Tr("Yekaterinburg"))
	checkZoneValid("Indian/Chagos", "UTC+06 "+Tr("Chagos"))
	checkZoneValid("Asia/Bangkok", "UTC+07 "+Tr("Bangkok"))
	checkZoneValid("Asia/Ho_Chi_Minh", "UTC+07 "+Tr("Ho Chi Minh"))
	checkZoneValid("Asia/Hovd", "UTC+07 "+Tr("Hovd"))
	checkZoneValid("Asia/Jakarta", "UTC+07 "+Tr("Jakarta"))
	checkZoneValid("Asia/Novokuznetsk", "UTC+07 "+Tr("Novokuznetsk"))
	checkZoneValid("Asia/Novosibirsk", "UTC+07 "+Tr("Novosibirsk"))
	checkZoneValid("Asia/Omsk", "UTC+07 "+Tr("Omsk"))
	checkZoneValid("Asia/Phnom_Penh", "UTC+07 "+Tr("Phnom Penh"))
	checkZoneValid("Asia/Pontianak", "UTC+07 "+Tr("Pontianak"))
	checkZoneValid("Asia/Saigon", "UTC+07 "+Tr("Saigon"))
	checkZoneValid("Asia/Vientiane", "UTC+07 "+Tr("Vientiane"))
	checkZoneValid("Indian/Christmas", "UTC+07 "+Tr("Christmas"))
	checkZoneValid("Asia/Brunei", "UTC+08 "+Tr("Brunei"))
	checkZoneValid("Asia/Calcutta", "UTC+08 "+Tr("Calcutta"))
	checkZoneValid("Asia/Chongqing", "UTC+08 "+Tr("Chongqing"))
	checkZoneValid("Asia/Harbin", "UTC+08 "+Tr("Harbin"))
	checkZoneValid("Asia/Hong_Kong", "UTC+08 "+Tr("Hong Kong"))
	checkZoneValid("Asia/Kashgar", "UTC+08 "+Tr("Kashgar"))
	checkZoneValid("Asia/Kathmandu", "UTC+08 "+Tr("Kathmandu"))
	checkZoneValid("Asia/Kuala_Lumpur", "UTC+08 "+Tr("Kuala Lumpur"))
	checkZoneValid("Asia/Kuching", "UTC+08 "+Tr("Kuching"))
	checkZoneValid("Asia/Macao", "UTC+08 "+Tr("Macao"))
	checkZoneValid("Asia/Makassar", "UTC+08 "+Tr("Makassar"))
	checkZoneValid("Asia/Manila", "UTC+08 "+Tr("Manila"))
	checkZoneValid("Asia/Shanghai", "UTC+08 "+Tr("Shanghai"))
	checkZoneValid("Asia/Beijing", "UTC+08 "+Tr("Beijing"))
	checkZoneValid("Asia/Singapore", "UTC+08 "+Tr("Singapore"))
	checkZoneValid("Asia/Taipei", "UTC+08 "+Tr("Taipei"))
	checkZoneValid("Asia/Ujung_Pandang", "UTC+08 "+Tr("Ujung Pandang"))
	checkZoneValid("Asia/Ulan_Bator", "UTC+08 "+Tr("Ulan Bator"))
	checkZoneValid("Asia/Urumqi", "UTC+08 "+Tr("Urumqi"))
	checkZoneValid("Australia/Perth", "UTC+08 "+Tr("Perth"))
	checkZoneValid("Australia/West", "UTC+08 "+Tr("Australia West"))
	checkZoneValid("Asia/Dili", "UTC+09 "+Tr("Dili"))
	checkZoneValid("Asia/Irkutsk", "UTC+09 "+Tr("Irkutsk"))
	checkZoneValid("Asia/Jayapura", "UTC+09 "+Tr("Jayapura"))
	checkZoneValid("Asia/Pyongyang", "UTC+09 "+Tr("Pyongyang"))
	checkZoneValid("Asia/Seoul", "UTC+09 "+Tr("Seoul"))
	checkZoneValid("Asia/Tokyo", "UTC+09 "+Tr("Tokyo"))
	checkZoneValid("Pacific/Palau", "UTC+09 "+Tr("Palau"))
	checkZoneValid("ROK", "UTC+09 "+Tr("ROK"))
	checkZoneValid("Antarctica/DumontDUrville", "UTC+10 "+Tr("DumontDUrville"))
	checkZoneValid("Asia/Yakutsk", "UTC+10 "+Tr("Yakutsk"))
	checkZoneValid("Australia/ACT", "UTC+10 "+Tr("ACT"))
	checkZoneValid("Australia/Adelaide", "UTC+10 "+Tr("Adelaide"))
	checkZoneValid("Australia/Broken_Hill", "UTC+10 "+Tr("Broken Hill"))
	checkZoneValid("Australia/Currie", "UTC+10 "+Tr("Currie"))
	checkZoneValid("Australia/Darwin", "UTC+10 "+Tr("Darwin"))
	checkZoneValid("Australia/Lindeman", "UTC+10 "+Tr("Lindeman"))
	checkZoneValid("Australia/Melbourne", "UTC+10 "+Tr("Melbourne"))
	checkZoneValid("Australia/North", "UTC+10 "+Tr("Australia North"))
	checkZoneValid("Australia/Queensland", "UTC+10 "+Tr("Queensland"))
	checkZoneValid("Australia/South", "UTC+10 "+Tr("Australia South"))
	checkZoneValid("Australia/Tasmania", "UTC+10 "+Tr("Tasmania"))
	checkZoneValid("Australia/Victoria", "UTC+10 "+Tr("Victoria"))
	checkZoneValid("Pacific/Chatham", "UTC+10 "+Tr("Chatham"))
	checkZoneValid("Pacific/Guam", "UTC+10 "+Tr("Guam"))
	checkZoneValid("Pacific/Port_Moresby", "UTC+10 "+Tr("Port Moresby"))
	checkZoneValid("Pacific/Saipan", "UTC+10 "+Tr("Saipan"))
	checkZoneValid("Pacific/Truk", "UTC+10 "+Tr("Truk"))
	checkZoneValid("Antarctica/Casey", "UTC+11 "+Tr("Casey"))
	checkZoneValid("Asia/Sakhalin", "UTC+11 "+Tr("Sakhalin"))
	checkZoneValid("Asia/Vladivostok", "UTC+11 "+Tr("Vladivostok"))
	checkZoneValid("Australia/Lord_Howe", "UTC+11 "+Tr("Lord Howe"))
	checkZoneValid("Pacific/Efate", "UTC+11 "+Tr("Efate"))
	checkZoneValid("Pacific/Efate", "UTC+11 "+Tr("Nouméa"))
	checkZoneValid("Pacific/Guadalcanal", "UTC+11 "+Tr("Guadalcanal"))
	checkZoneValid("Pacific/Kosrae", "UTC+11 "+Tr("Kosrae"))
	checkZoneValid("Pacific/Norfolk", "UTC+11 "+Tr("Norfolk"))
	checkZoneValid("Pacific/Ponape", "UTC+11 "+Tr("Ponape"))
	checkZoneValid("Antarctica/McMurdo", "UTC+12 "+Tr("McMurdo"))
	checkZoneValid("Antarctica/South_Pole", "UTC+12 "+Tr("South Pole"))
	checkZoneValid("Asia/Anadyr", "UTC+12 "+Tr("Anadyr"))
	checkZoneValid("Asia/Kabul", "UTC+12 "+Tr("Kabul"))
	checkZoneValid("Asia/Magadan", "UTC+12 "+Tr("Magadan"))
	checkZoneValid("NZ", "UTC+12 "+Tr("Wellington"))
	checkZoneValid("Pacific/Auckland", "UTC+12 "+Tr("Auckland"))
	checkZoneValid("Pacific/Fiji", "UTC+12 "+Tr("Fiji"))
	checkZoneValid("Pacific/Funafuti", "UTC+12 "+Tr("Funafuti"))
	checkZoneValid("Pacific/Kwajalein", "UTC+12 "+Tr("Kwajalein"))
	checkZoneValid("Pacific/Nauru", "UTC+12 "+Tr("Nauru"))
	checkZoneValid("Pacific/Tarawa", "UTC+12 "+Tr("Tarawa"))
	checkZoneValid("Pacific/Wake", "UTC+12 "+Tr("Wake"))
	checkZoneValid("Pacific/Wallis", "UTC+12 "+Tr("Wallis"))
	checkZoneValid("Pacific/Enderbury", "UTC+13 "+Tr("Enderbury"))
	checkZoneValid("Pacific/Fakaofo", "UTC+13 "+Tr("Fakaofo"))
	checkZoneValid("Pacific/Tongatapu", "UTC+13 "+Tr("Tongatapu"))
}
