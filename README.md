# OYMeals

Fetches todays (tommorrows if time>1700) meals for a few restaurants at the
campus of university of Oulu.

## Restaurants

- Mara (Oulu university of applied sciences)
- Kerttu
- Voltti
- Lipasto & Juliana (common menu)

## Installation

```sh
go install github.com/adventune/oymeals
```

## Usage

The program can be used both as a webserver and a SSG. Run with:

```sh
oymeals
```

| Env     | Description                                        | Default        |
| ------- | -------------------------------------------------- | -------------- |
| STATIC  | If set, generates a static HTML with the menu.     | 0              |
| OUTFILE | Sets the file name and path of the generated file. | `./meals.html` |
| TITLE   | Sets a custom title for the page.                  | OYMeals        |
