# Notas sobre Golang

Notas gerais de artigos, postagens e anotações sobre a linguagem.

## Goroutines

Goroutines são como threads "virtuais", porém:

- Começam com apenas 2kb de memória, crescendo conforme necessário
- São gerenciadas pelo runtime do Go, sem depender diretamente do sistema operacional
- Permitem criar milhares (ou até milhões) de tarefas concorrentes sem sobrecarregar a memória

Sendo o modelo **M:P:G** responsável pelo gerenciamento das Gourotines pelo runtime, combinando:

- M (Threads reais): criadas pelo sistema operacional
- P (Processors): Responsáveis por gerenciar as Gourotines
- G (Goroutines): Tarefas concorrentes atribuídas a Processors

O cérebro desse runtime se chama scheduler e ele é como um mini sistema operacional que atribui tarefas para que os Processors não fiquem ociosos, redistribui Goroutines entre Processors com Work Stealing, evitando gargalos e gerencia chamadas de I/O liberando threads reais para outras tarefas enquanto espera.

Go gerencia a memória com Goroutines tendo stacks independentes que crescem conforme necessário. O Garbage Collector garante que a memória seja liberada quando não é mais usada e o runtime ajusta dinamicamente a quantidade de threads reais, balanceando desempenho e consumo de recursos.
