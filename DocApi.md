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