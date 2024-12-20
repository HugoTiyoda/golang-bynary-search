# Comparação de Algoritmos de Busca em Go

Este projeto implementa e compara o desempenho entre dois algoritmos fundamentais de busca: busca linear (pesquisa sequencial) e busca binária.

##  Sobre o Projeto

O objetivo é demonstrar a diferença de performance entre os dois algoritmos de busca mais comuns, utilizando Go e seus recursos nativos de benchmark.

### Implementações

1. **Busca Linear (O(n))**
   - Percorre sequencialmente cada elemento do array
   - Complexidade: O(n)
   - Vantagem: Funciona com arrays não ordenados
   - Desvantagem: Menos eficiente para grandes conjuntos de dados

2. **Busca Binária (O(log n))**
   - Divide o espaço de busca pela metade a cada iteração
   - Complexidade: O(log n)
   - Vantagem: Muito mais eficiente para grandes conjuntos
   - Desvantagem: Requer array ordenado

##  Como Executar

1. Clone o repositório
2. Navegue até o diretório do projeto
3. Execute os testes e benchmarks:

```bash
go test -bench=.
```

##  Resultados dos Benchmarks

Executado em um processador Intel Core i9-12900H:

### Busca Linear
```
BenchmarkLinearSearch-14    40006184    30.90 ns/op
```
- 40.006.184 iterações realizadas
- Média de 30.90 nanosegundos por operação
- Zero alocações de memória

### Busca Binária
```
BenchmarkBinarySearch-14    23348906    5.062 ns/op
```
- 23.348.906 iterações realizadas
- Média de 5.062 nanosegundos por operação
- Zero alocações de memória

### Análise dos Resultados

A busca binária demonstrou ser aproximadamente 6 vezes mais rápida que a busca linear no cenário testado (array de 100 elementos). Este resultado está alinhado com a complexidade teórica dos algoritmos:

- Busca Linear: O(n) - até 100 comparações no pior caso
- Busca Binária: O(log n) - aproximadamente 7 comparações no pior caso (log₂ 100 ≈ 6.64)

##  Detalhes Técnicos

- Linguagem: Go 1.22.4
- Ambiente: Windows
- Arquitetura: AMD64
- CPU: 12th Gen Intel(R) Core(TM) i9-12900H

##  Estrutura do Código

O projeto contém um único arquivo `search_test.go` com:
- Implementações dos algoritmos de busca
- Função auxiliar para geração de arrays
- Testes unitários
- Benchmarks para casos de pior cenário
- Função main para demonstração

##  Considerações

1. Os benchmarks foram executados no pior caso (buscando o último elemento)
2. O array utilizado tem 100 elementos
3. Todas as medições apresentaram zero alocações de memória
4. Os resultados podem variar dependendo do hardware utilizado

##