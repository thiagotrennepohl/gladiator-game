package mocks

type OpponentMock struct {
	name string
	str  float32
	sta  float32
	dex  float32
}

func (mock *OpponentMock) GetAttackPoints() float32 {
	return mock.str
}

func (mock *OpponentMock) GetDefensePoints() float32 {
	return mock.sta
}

func (mock *OpponentMock) GetName() string {
	return mock.name
}

func (mock *OpponentMock) GetEvasionRate() float32 {
	//Create rule for evasion rate based on luck and dextery
	return mock.dex
}

func (mock *OpponentMock) GetCriticalChance() float32 {
	return mock.dex
}

func (mock *OpponentMock) SetName(name string) {
	mock.name = name
}

func (mock *OpponentMock) SetAttackPoints(ap float32) {
	mock.str = ap
}
