package moto

//Структуры должны содержать марку авто, год выпуска, объем багажника/кузова, запущен ли двигатель,
// открыты ли окна, насколько заполнен объем багажника.

//Moto -
type Moto struct {
	Brand            string
	Model            string
	ProdYear         int
	BootVolumeLiters int
	IsEngineOn       bool
	IsWindowsOpen    bool
	BootEmptyPercent float32
}

//Truck -
type Truck struct {
	Chassis   Moto
	BootClass int
}

//BootClass consts for Trunk

//BootClassUncovered -
const BootClassUncovered int = 0

//BootClassCovered -
const BootClassCovered int = 1

//BootClassRefrigerator -
const BootClassRefrigerator int = 2

//Car -
type Car struct {
	Chassis      Moto
	ChassisClass int
}

//ChassisClass consts for Trunk

//ChassisClassSedan -
const ChassisClassSedan int = 0

//ChassisClassHatchback -
const ChassisClassHatchback int = 1

//ChassisClassStationWagon -
const ChassisClassStationWagon int = 2

//ChassisClassPickup -
const ChassisClassPickup int = 3
