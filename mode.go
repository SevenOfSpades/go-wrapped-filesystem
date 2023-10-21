package filesystem

import "os"

const (
	ModeModifierRead    ModeModifier = 04
	ModeModifierWrite   ModeModifier = 02
	ModeModifierExecute ModeModifier = 01

	ModeAdjustmentUser   ModeAdjustment = 6
	ModeAdjustmentGroup  ModeAdjustment = 3
	ModeAdjustmentOthers ModeAdjustment = 0

	ModeUserRead             = Mode(ModeModifierRead << ModeAdjustmentUser)
	ModeUserWrite            = Mode(ModeModifierWrite << ModeAdjustmentUser)
	ModeUserExecute          = Mode(ModeModifierExecute << ModeAdjustmentUser)
	ModeUserReadWrite        = ModeUserRead | ModeUserWrite
	ModeUserReadWriteExecute = ModeUserReadWrite | ModeUserExecute

	ModeGroupRead             = Mode(ModeModifierRead << ModeAdjustmentGroup)
	ModeGroupWrite            = Mode(ModeModifierWrite << ModeAdjustmentGroup)
	ModeGroupExecute          = Mode(ModeModifierExecute << ModeAdjustmentGroup)
	ModeGroupReadWrite        = ModeGroupRead | ModeGroupWrite
	ModeGroupReadWriteExecute = ModeGroupReadWrite | ModeGroupExecute

	ModeOthersRead             = Mode(ModeModifierRead << ModeAdjustmentOthers)
	ModeOthersWrite            = Mode(ModeModifierWrite << ModeAdjustmentOthers)
	ModeOthersExecute          = Mode(ModeModifierExecute << ModeAdjustmentOthers)
	ModeOthersReadWrite        = ModeOthersRead | ModeOthersWrite
	ModeOthersReadWriteExecute = ModeOthersReadWrite | ModeOthersExecute

	ModeAllRead             = ModeUserRead | ModeGroupRead | ModeOthersRead
	ModeAllWrite            = ModeUserWrite | ModeGroupWrite | ModeOthersWrite
	ModeAllExecute          = ModeUserExecute | ModeGroupExecute | ModeOthersExecute
	ModeAllReadWrite        = ModeAllRead | ModeAllWrite
	ModeAllReadWriteExecute = ModeAllReadWrite | ModeAllExecute
)

type (
	Mode           uint32
	ModeAdjustment uint32
	ModeModifier   uint32
)

func (m Mode) asFileMode() os.FileMode {
	return os.FileMode(m)
}
