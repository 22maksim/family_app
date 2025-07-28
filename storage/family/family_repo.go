package family

type FamilyRepository interface {
	CreateFamily() string
}

type FamilyRepositoryImpl struct {
	Name string
}

func (fr *FamilyRepositoryImpl) CreateFamily() string {
	return "12345"
}

func NewFamilyRepository() FamilyRepository {
	return &FamilyRepositoryImpl{
		Name: "MyMockDb",
	}
}
