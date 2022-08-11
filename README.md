# Go Todo
![hero](https://i.imgur.com/e7tOUCO.png)

## Introduction

A simple todolist application written in Go 

## Requirements
* MySQL installed
* Go installed

## Installation

* This repo is fork from repo:
git clone https://github.com/ichtrojan/go-todo.git

* Clone this repo 

```bash
git clone https://github.com/ryosuryo/go-todo.git
```

* Change Directory

```bash
cd go-todo
```

* Initiate `.env` file

```bash
cp .env.example .env
```

* Modify `.env` file with your correct database credentials and desired Port

## Usage

To run this application, execute:

```bash
go run main.go
```

You should be able to access this application at `http://127.0.0.1:6060`

>**NOTE**<br>
>If you modified the port in the `.env` file, you should access the application for the port you set

## Conclusion 

This Project is an example to teach CRUD using the default `database/sql` package and how to serve html templates properly.

If you have anything to add to this, please send in a PR as it will no longer be actively maintained by [me](https://github.com/ichtrojan).

Live Preview : http://godamar.ioiapps.com/
If the live preview site can't access fell free to contact me from [WhatsApp](https://api.whatsapp.com/send/?phone=6285799990062&text&type=phone_number&app_absent=0)