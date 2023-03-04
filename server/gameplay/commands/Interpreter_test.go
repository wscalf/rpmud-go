package commands

import (
	"rpmud/server/gameplay/commands/builtin"
	"rpmud/server/gameplay/commands/parameters"
	"rpmud/server/gameplay/world"
	"rpmud/server/infrastructure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkInterpreterParsingComplexCommand(b *testing.B) {
	p, r := stageTestData()

	in := NewInterpreter()
	in.Register("p", TestPage{})

	for i := 0; i < b.N; i++ {
		_, _ = in.Prepare(`p Rumil "Sorkhild the Undaunted"=Hey, guys!`, p, r)
	}
}

func TestInterpreterCreatesLinkJobFromLinkName(t *testing.T) {
	p, r := stageTestData()

	i := NewInterpreter()

	j, err := i.Prepare("link", p, r)
	assert.NoError(t, err)

	if job, ok := j.(LinkJob); ok {
		assert.Equal(t, "link", job.link.Command)
	} else {
		assert.Fail(t, "Job did not convert to LinkJob")
	}
}

func TestInterpreterCreatesCmdJobFromParameterlessCommand(t *testing.T) {
	p, r := stageTestData()

	i := NewInterpreter()

	i.Register("noparams", Parameterless{})

	j, err := i.Prepare("noparams", p, r)
	assert.NoError(t, err)

	if job, ok := j.(CmdJob); ok {
		_, ok = job.command.(Parameterless)
		assert.True(t, ok, "Concrete command was not Parameterless type")
		assert.Equal(t, 0, len(job.params))
	} else {
		assert.Fail(t, "Job did not convert to CmdJob")
	}
}

func TestInterpreterCreatesCmdJobFromComplexCommand(t *testing.T) {
	p, r := stageTestData()

	i := NewInterpreter()

	i.Register("p", TestPage{})

	j, err := i.Prepare(`p Rumil "Sorkhild the Undaunted"=Hey, guys!`, p, r)
	assert.NoError(t, err)

	if job, ok := j.(CmdJob); ok {
		_, ok = job.command.(TestPage)
		assert.True(t, ok, "Concrete command was not TestPage type")
		assert.Equal(t, len(job.params), 3)

		assert.Equal(t, []string{"Rumil", "Sorkhild the Undaunted"}, job.params["names"].Multiple())
		assert.Equal(t, "=", job.params["="].Single())
		assert.Equal(t, "Hey, guys!", job.params["message"].Single())
	} else {
		assert.Fail(t, "Job did not convert to CmdJob")
	}
}

func TestInterpreterReturnsInvalidInputForUnrecognizedCommand(t *testing.T) {
	p, r := stageTestData()

	i := NewInterpreter()

	_, err := i.Prepare("made-up command", p, r)

	assert.ErrorIs(t, err, ErrInvalidInput)
}

func TestInterpreterOmitsMissingOptionalParameters(t *testing.T) {
	p, r := stageTestData()

	i := NewInterpreter()

	j, err := i.Prepare("look", p, r)
	assert.NoError(t, err)

	if job, ok := j.(CmdJob); ok {
		_, ok = job.command.(*builtin.Look)
		assert.True(t, ok, "Concrete command was not Look")
		assert.Equal(t, 0, len(job.params))
	} else {
		assert.Fail(t, "Job did not convert to CmdJob")
	}
}

func TestInterpreterCapturesOptionalParameters(t *testing.T) {
	p, r := stageTestData()

	i := NewInterpreter()

	j, err := i.Prepare("look obj", p, r)
	assert.NoError(t, err)

	if job, ok := j.(CmdJob); ok {
		_, ok = job.command.(*builtin.Look)
		assert.True(t, ok, "Concrete command was not Look")
		assert.Equal(t, 1, len(job.params))
		assert.Equal(t, "obj", job.params["at"].Single())
	} else {
		assert.Fail(t, "Job did not convert to CmdJob")
	}
}

func TestInterpreterReturnsInvalidInputForMissingRequiredParam(t *testing.T) {
	p, r := stageTestData()

	i := NewInterpreter()

	i.Register("p", TestPage{})

	_, err := i.Prepare(`p Rumil Hey, guys!`, p, r)

	assert.ErrorIs(t, err, ErrInvalidInput)
}

func stageTestData() (player *world.Player, room *world.Room) {
	room = world.NewRoom("start", "A starting place.")
	other := world.NewRoom("other", "Another place")
	room.LinkTo(other, "link", "Test Link", "The test link.")

	player = world.NewPlayer(infrastructure.StubClient{}, func(p *world.Player, r *world.Room, s string) {}, "player")
	player.Enter(room)

	return
}

type Parameterless struct {
}

func (pc Parameterless) Params() []parameters.Parameter {
	return []parameters.Parameter{}
}

func (pc Parameterless) Exec(p *world.Player, r *world.Room, params map[string]parameters.Value) {

}

type TestPage struct {
}

func (pc TestPage) Params() []parameters.Parameter {
	return []parameters.Parameter{
		parameters.NewNameGroup("names", true),
		parameters.NewDelimiter("="),
		parameters.NewFreeText("message", true),
	}
}

func (pc TestPage) Exec(p *world.Player, r *world.Room, params map[string]parameters.Value) {

}
