package main

import (
	"bufio" // Biblioteca para leitura de entradas de forma eficiente
	"fmt"   // Biblioteca para formatação e impressão de dados
	"os"    // Biblioteca para interagir com o sistema operacional (e.g., manipulação de arquivos)
)

func main() {
	// Cria um mapa para armazenar a contagem de cada linha (chave: string, valor: int)
	counts := make(map[string]int)

	// Recebe os argumentos passados na linha de comando, ignorando o primeiro (que é o nome do programa)
	files := os.Args[1:]

	// Verifica se não foram passados arquivos como argumento
	if len(files) == 0 {
		// Caso não tenha arquivos, lê a entrada padrão (por exemplo, o que for digitado no terminal)
		countLines(os.Stdin, counts)
	} else {
		// Caso tenha arquivos, percorre a lista de arquivos fornecida
		for _, arg := range files {
			// Tenta abrir o arquivo correspondente
			f, err := os.Open(arg)
			if err != nil {
				// Em caso de erro ao abrir o arquivo, imprime uma mensagem de erro para o stderr e continua com o próximo arquivo
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			// Conta as linhas do arquivo aberto
			countLines(f, counts)
			// Fecha o arquivo após a leitura para liberar recursos do sistema
			f.Close()
		}
	}

	// Itera sobre o mapa 'counts' que contém as linhas e suas contagens
	for line, n := range counts {
		// Imprime apenas as linhas que aparecem mais de uma vez
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Printf(files[0]) // Esta linha imprime o nome do arquivo que vem como argumento
		}
	}
}

// Função que conta as linhas em um arquivo e armazena as contagens no mapa
func countLines(f *os.File, counts map[string]int) {
	// Cria um scanner para ler o arquivo linha por linha
	input := bufio.NewScanner(f)
	// Lê cada linha e a armazena no mapa 'counts', incrementando o contador para cada ocorrência
	for input.Scan() {
		counts[input.Text()]++
	}
	// Nota: Ignora erros que podem ocorrer durante a leitura (como falhas no scanner)
}
