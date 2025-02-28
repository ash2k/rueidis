// Code generated DO NOT EDIT

package cmds

import "testing"

func bf0(s Builder) {
	s.BfAdd().Key("1").Item("1").Build()
	s.BfCard().Key("1").Build()
	s.BfExists().Key("1").Item("1").Build()
	s.BfExists().Key("1").Item("1").Cache()
	s.BfInfo().Key("1").Capacity().Build()
	s.BfInfo().Key("1").Capacity().Cache()
	s.BfInfo().Key("1").Size().Build()
	s.BfInfo().Key("1").Size().Cache()
	s.BfInfo().Key("1").Filters().Build()
	s.BfInfo().Key("1").Filters().Cache()
	s.BfInfo().Key("1").Items().Build()
	s.BfInfo().Key("1").Items().Cache()
	s.BfInfo().Key("1").Expansion().Build()
	s.BfInfo().Key("1").Expansion().Cache()
	s.BfInfo().Key("1").Build()
	s.BfInfo().Key("1").Cache()
	s.BfInsert().Key("1").Capacity(1).Error(1).Expansion(1).Nocreate().Nonscaling().Items().Item("1").Item("1").Build()
	s.BfInsert().Key("1").Capacity(1).Error(1).Expansion(1).Nocreate().Items().Item("1").Item("1").Build()
	s.BfInsert().Key("1").Capacity(1).Error(1).Expansion(1).Nonscaling().Items().Item("1").Item("1").Build()
	s.BfInsert().Key("1").Capacity(1).Error(1).Expansion(1).Items().Item("1").Item("1").Build()
	s.BfInsert().Key("1").Capacity(1).Error(1).Nocreate().Nonscaling().Items().Item("1").Item("1").Build()
	s.BfInsert().Key("1").Capacity(1).Error(1).Nonscaling().Items().Item("1").Item("1").Build()
	s.BfInsert().Key("1").Capacity(1).Error(1).Items().Item("1").Item("1").Build()
	s.BfInsert().Key("1").Capacity(1).Expansion(1).Nocreate().Nonscaling().Items().Item("1").Item("1").Build()
	s.BfInsert().Key("1").Capacity(1).Nocreate().Nonscaling().Items().Item("1").Item("1").Build()
	s.BfInsert().Key("1").Capacity(1).Nonscaling().Items().Item("1").Item("1").Build()
	s.BfInsert().Key("1").Capacity(1).Items().Item("1").Item("1").Build()
	s.BfInsert().Key("1").Error(1).Expansion(1).Nocreate().Nonscaling().Items().Item("1").Item("1").Build()
	s.BfInsert().Key("1").Expansion(1).Nocreate().Nonscaling().Items().Item("1").Item("1").Build()
	s.BfInsert().Key("1").Nocreate().Nonscaling().Items().Item("1").Item("1").Build()
	s.BfInsert().Key("1").Nonscaling().Items().Item("1").Item("1").Build()
	s.BfInsert().Key("1").Items().Item("1").Item("1").Build()
	s.BfLoadchunk().Key("1").Iterator(1).Data("1").Build()
	s.BfMadd().Key("1").Item("1").Item("1").Build()
	s.BfMexists().Key("1").Item("1").Item("1").Build()
	s.BfReserve().Key("1").ErrorRate(1).Capacity(1).Expansion(1).Nonscaling().Build()
	s.BfReserve().Key("1").ErrorRate(1).Capacity(1).Expansion(1).Build()
	s.BfReserve().Key("1").ErrorRate(1).Capacity(1).Nonscaling().Build()
	s.BfReserve().Key("1").ErrorRate(1).Capacity(1).Build()
	s.BfScandump().Key("1").Iterator(1).Build()
}

func TestCommand_InitSlot_bf(t *testing.T) {
	var s = NewBuilder(InitSlot)
	t.Run("0", func(t *testing.T) { bf0(s) })
}

func TestCommand_NoSlot_bf(t *testing.T) {
	var s = NewBuilder(NoSlot)
	t.Run("0", func(t *testing.T) { bf0(s) })
}
