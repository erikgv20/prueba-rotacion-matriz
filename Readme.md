# PRUEBA GO V1
1. Validación matriz NxN
2. Rotación horario de una matriz multi-dimensional
3. INPUT/OUTPUT JSON
4. Test Postman
5. Puerto 10000 (localhost:10000) 

## Run
```bash
$ go run main.go
```


## Input:
```json
{
    "dataform": [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9]
    ]   
}
```


## Output:
```json
{
    "ArrayValido": "Array válido",
    "ArrayOriginal": [
        [
            1,
            2,
            3
        ],
        [
            4,
            5,
            6
        ],
        [
            7,
            8,
            9
        ]
    ],
    "ArrayProcesado": [
        [
            7,
            4,
            1
        ],
        [
            8,
            5,
            2
        ],
        [
            9,
            6,
            3
        ]
    ]
}
```
