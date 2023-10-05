# API


## /api/data/names GET

Donnes les noms des datasets disponibles

### Response

```json
[
  "les-arbres-plantes",
  "espaces_verts"
]
```
si erreur:
```json
{
  "error": "message d'erreur"
}
```

## /api/data/category GET

Donnes les datasets par catégories

si erreur:
```json
{
"error": "message d'erreur"
}
```

## /data/category/{category}

Donnes les datasets de la catégorie {category}

si erreur:
```json
{
"error": "message d'erreur"
}
```

## /api/data/name/{name} GET

Donnes les données du dataset {name}

si erreur:
```json
{
  "error": "message d'erreur"
}
```

## /api/data/name/{name}/fields GET

Donnes les champs du dataset {name}

paramètres optionnels:

- rmId : bool (default: false) : si true, supprime le champ _id
- rmMulFields : bool (default: false) : si true, supprime les champs qui sont des tableaux


si erreur:
```json
{
  "error": "message d'erreur"
}
```

## /api/data/name/{name}/filter GET

require query params:
- q : query string

q est un json qui sert de filtre avec ce format:
```json
{
  "field1": "value1",
  "field2": "value2",
  "field3": "value3"
}
```

Donnes une donnée du dataset {name} filtrée par q

si erreur:
```json
{
  "error": "message d'erreur"
}
```

## /api/data/name/{name}/filters GET

require query params:
- q : query string

q est un json qui sert de filtre avec ce format:
```json
{
  "field1": "value1",
  "field2": "value2",
  "field3": "value3"
}
```

Donnes les données du dataset {name} filtrées par q

si erreur:
```json
{
  "error": "message d'erreur"
}
```


## /api/data/name/{name}/field/{field}/value/{value} GET

Donnes les données du dataset {name} où le champ {field} est égal à {value}

si erreur:
```json
{
  "error": "message d'erreur"
}
```

## /api/graph/{name} GET

Donnes le type de graph du dataset {name}

si erreur:
```json
{
  "error": "message d'erreur"
}
```

## /api/graph GET

Donnes les types de graph des datasets