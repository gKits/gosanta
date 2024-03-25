package gosanta

import (
	"fmt"
	"math/rand/v2"
	"net/mail"
	"strings"
	"text/tabwriter"
)

type Participant struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Shuffler struct {
	participants []Participant
	names        map[string]struct{}
}

func NewShuffler() *Shuffler {
	return &Shuffler{
		participants: make([]Participant, 0),
		names:        make(map[string]struct{}),
	}
}

func (shuff *Shuffler) AddParticipant(part Participant) error {
	if _, ok := shuff.names[part.Name]; ok {
		return fmt.Errorf("name %s is mapped more than onces", part.Name)
	}
	if _, err := mail.ParseAddress(part.Email); err != nil {
		return fmt.Errorf("%s is not a valid e-mail address", part.Email)
	}

	shuff.names[part.Name] = struct{}{}
	shuff.participants = append(shuff.participants, part)
	return nil
}

func (shuff *Shuffler) Shuffle() map[Participant]string {
	rand.Shuffle(len(shuff.participants), func(i, j int) {
		shuff.participants[i], shuff.participants[j] = shuff.participants[j], shuff.participants[i]
	})

	res := make(map[Participant]string)
	for i, part := range shuff.participants {
		res[part] = shuff.participants[(i+1)%len(shuff.participants)].Name
	}

	return res
}

func (shuff *Shuffler) String() string {
	builder := new(strings.Builder)
	w := tabwriter.NewWriter(builder, 1, 1, 1, ' ', 0)
	for _, part := range shuff.participants {
		fmt.Fprintln(w, part.Name, "\t", part.Email)
	}
	w.Flush()
	return builder.String()
}
