package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "modernc.org/sqlite"
)

func TestSelectClientWhenOk(t *testing.T) {
	//подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		require.NoError(t, err)
	}
	defer db.Close()

	clientID := 1
	//получаем объект клиента
	cl, err := selectClient(db, clientID)
	//проверка на ошибку
	require.NoError(t, err)
	//сравниваем что поле ID объекта Client совпадает с индификатором переменной
	assert.Equal(t, clientID, cl.ID)
	//проверяем что остальные поля не пустые
	assert.NotEmpty(t, cl.FIO)
	assert.NotEmpty(t, cl.Login)
	assert.NotEmpty(t, cl.Birthday)
	assert.NotEmpty(t, cl.Email)
}

func TestSelectClientWhenNoClient(t *testing.T) {
	//подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		require.NoError(t, err)
	}
	defer db.Close()

	clientID := -1
	//получаем объект клиента
	cl, err := selectClient(db, clientID)
	//сравнием что функция вернула ошику
	require.Equal(t, sql.ErrNoRows, err)
	//проверяем что все поля пустые
	assert.Empty(t, cl.ID)
	assert.Empty(t, cl.FIO)
	assert.Empty(t, cl.Login)
	assert.Empty(t, cl.Birthday)
	assert.Empty(t, cl.Email)
}

func TestInsertClientThenSelectAndCheck(t *testing.T) {
	//подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		require.NoError(t, err)
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}
	//добавляем запись в таблицу
	cl.ID, err = insertClient(db, cl)
	//проверка на ошибку
	require.NoError(t, err)
	//проверяем что поле не пустое
	require.NotEmpty(t, cl.ID)
	//получаем объект Client по индефикатору
	stored, err := selectClient(db, cl.ID)
	//проверяем на ошибку
	require.NoError(t, err)
	//равниваем значения полей
	assert.Equal(t, cl, stored)
}

func TestInsertClientDeleteClientThenCheck(t *testing.T) {
	//подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		require.NoError(t, err)
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}
	//добавляем запись в таблицу
	id, err := insertClient(db, cl)
	//проверка на ошибку
	require.NoError(t, err)
	//проверяем что поле не пустое
	require.NotEmpty(t, id)
	//получаем объект клиента
	_, err = selectClient(db, id)
	//проверяем на ошибку
	require.NoError(t, err)
	//удаляем запись
	err = deleteClient(db, id)
	//проверяем на ошибку
	require.NoError(t, err)
	//получаем объект клиента
	_, err = selectClient(db, id)
	//сравниваем на ошибку
	require.Equal(t, sql.ErrNoRows, err)
}
