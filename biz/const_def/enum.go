package const_def

type RelationEnum int32

const (
	RelationEnumFollow       RelationEnum = 1
	RelationEnumCancelFollow RelationEnum = -1
)

var (
	RelationEnumStrMap = map[RelationEnum]string{
		RelationEnumCancelFollow: "取消关注",
		RelationEnumFollow:       "关注",
	}
)

var (
	RelationEnumEngMap = map[RelationEnum]string{
		RelationEnumFollow: "follow",
	}
)

type ActionEnum int32

const (
	ActionEnumCollectGame       ActionEnum = 1
	ActionEnumReserveGame       ActionEnum = 2
	ActionEnumCanCelCollectGame ActionEnum = -1
	ActionEnumCancelReserveGame ActionEnum = -2
)

var ()
