package config

type DBconf struct {
	USER string
	DATABASE string
	PASSWORD string
	HOST string
	PORT string
}

type Connect_data struct {
	DB DBconf
}

func Config_data() Connect_data {
	var all_data Connect_data
	
	database := DBconf{
		USER: "jphacks",
		DATABASE: "jphacks",
		PASSWORD: "jphacksdb",
		HOST: "jphacksdb.cc47ubzedfw4.ap-northeast-1.rds.amazonaws.com",
		PORT: "5432",
	}

	all_data.DB = database

	return all_data
}

