# Integra-ocontinua

PRINCIPAIS PROCESSOS DA INTEGRAÇÃO CONTINUA 
-Execução de testes
-linter
-verificações de segurança
-geração de artefatos prontos para o processo de deploy
-Identifixação da próxima versão a ser gerada no software
-Geração de tags e releases

Utilizar a intehgração continua é bom para evitar debitos de aplicação, pois não vamos empurrando o problema com a barriga

STATUS CHECK 

bom o status chek é a garantia que um pull request não podera ser mergeada ao repositorio sem antes ter passado pelo processo de CI

FERRAMENTAS POPULARES: 

-jenkins 
-github actions 
-circle ci 
-aws code build 
-azure devops
-google cloud build
-gitlab pipelines/ci


O QUE É O GITHUB ACTIONS: 

Bom o github actions é um automarizador de workflow de desenvolvimento de software. Ele utiliza os principais eventos gerados pelo github para que possamos executar tarefas dos mais variados tipos, incluindo processos de CI.

DINAMICA DO GITHUB ACTIONS: 
O fluxo começa apartir de um workflow. O que seria um workfow?
bom um workflow é um conjunto de processos definidos por você. um exemplo seria fazer o build + rodar testes da aplicação 
algo interessante é que podemos ter mais de um workflow por repositorio
os wrokflows são definidos por arquivos .yml em github/workslows
ele pode possuir um ou mais jobs 
podemos configurar quando esse workflow é utilizado podemos botar para rodar por base em agendamento, por exemplo todos os dias a meia noite ele roda ou rodar sempre que fizerem um pull request
