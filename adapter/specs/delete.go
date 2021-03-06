package specs

import (
	"testing"

	"github.com/Fs02/grimoire"
	"github.com/Fs02/grimoire/adapter/sql"
	"github.com/Fs02/grimoire/c"
	"github.com/stretchr/testify/assert"
)

// Delete tests delete specifications.
func Delete(t *testing.T, repo grimoire.Repo) {
	record := User{Name: "delete", Age: 100}
	assert.Nil(t, repo.From(users).Save(&record))
	assert.Nil(t, repo.From(users).Save(&User{Name: "delete", Age: 100}))
	assert.Nil(t, repo.From(users).Save(&User{Name: "delete", Age: 100}))
	assert.Nil(t, repo.From(users).Save(&User{Name: "other delete", Age: 110}))

	tests := []grimoire.Query{
		repo.From(users).Find(record.ID),
		repo.From(users).Where(c.Eq(name, "delete")),
		repo.From(users).Where(c.Eq(name, "other delete"), c.Gt(age, 100)),
	}

	for _, query := range tests {
		statement, _ := sql.NewBuilder("?", false).Delete(query.Collection, query.Condition)
		t.Run("Delete|"+statement, func(t *testing.T) {
			var result []User
			assert.Nil(t, query.All(&result))
			assert.NotEqual(t, 0, len(result))

			assert.Nil(t, query.Delete())

			assert.Nil(t, query.All(&result))
			assert.Equal(t, 0, len(result))
		})
	}
}
