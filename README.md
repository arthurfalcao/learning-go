# learning-go

## Instalation

* Download: ``https://golang.org/dl/``
* Extract: ``sudo tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz``
* Add to the path: open ``sudo vim /etc/profile`` and add ``export PATH=$PATH:/usr/local/go/bin``

## Run

```zsh
  go run main.go
```

## commands

* **reflect.TypeOf(variable)** retorna o tipo
* **array := []string{"string"}** é um slice, que pode modificar o tamanho do array
* **var array = [4]string** é um array, com tamanho fixo
* **len(array)** retorna o tamanho do array
* **cap(array)** retorna a capaciadade do array
* **array = append(array, "string")** adiciona um novo item ao array e retorna