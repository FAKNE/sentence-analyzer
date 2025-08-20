# Go Microservice com Pipeline CI/CD para K3s em EC2

Este projeto foi desenvolvido como parte de uma avaliação técnica para a posição de **Platform Engineer**.  
Ele consiste em um **microserviço escrito em Go** que expõe uma API REST capaz de analisar uma frase e retornar o número de palavras, vogais e consoantes.  
A grande ênfase do projeto está na **automação de ponta a ponta**, cobrindo desde a **infraestrutura em nuvem** até a **entrega contínua via CI/CD**.

---

## 🎯 Objetivo da Solução

O objetivo é demonstrar competências em:

- **Infraestrutura como Código (IaC)** usando **Terraform**
- **Configuração automática** de servidores com **Ansible**
- **Containerização** com **Docker**
- **Orquestração em Kubernetes (K3s)**
- **Integração Contínua e Entrega Contínua (CI/CD)** usando **GitHub Actions**
- **Boas práticas de DevSecOps** (testes, segurança, monitoramento)

---

## 📂 Estrutura do Projeto

A seguir uma descrição detalhada de cada diretório e arquivo presente no repositório:

### `app/`

Contém o código-fonte do microserviço em Go.

- **main.go** → ponto de entrada da aplicação, define o endpoint REST.
- **handlers.go** → contém a lógica de processamento da frase (contagem de palavras, vogais, consoantes).
- **tests/** → testes unitários que validam o comportamento da aplicação.

### `terraform/`

Scripts de **Terraform** que definem a infraestrutura na AWS.

- **main.tf** → cria a instância EC2, define Security Groups e configura chaves SSH.
- **variables.tf** → parametriza recursos como tipo de instância e região.
- **outputs.tf** → exporta informações úteis (ex.: IP da instância) para o pipeline.

### `ansible/`

Playbooks do **Ansible** responsáveis pela configuração do servidor.

- **playbook.yml** → instala pacotes necessários, configura Docker e instala o K3s.
- **hosts** → inventário com o IP da instância EC2 provisionada.

### `deploy/`

Manifests Kubernetes para rodar a aplicação dentro do cluster K3s.

- **deployment.yaml** → define o número de réplicas e o container a ser executado.
- **service.yaml** → expõe o microserviço via NodePort para acesso externo.

### `Dockerfile`

Arquivo responsável por criar a imagem da aplicação Go.  
Foi construído como **multi-stage build**, para reduzir o tamanho final da imagem.

### `.github/workflows/ci-cd.yml`

Arquivo de configuração do pipeline **GitHub Actions**.  
Ele orquestra todas as etapas: provisionamento, configuração, build, push e deploy.

### `openapi.yaml`

Especificação **OpenAPI** (Swagger) da API, descrevendo o endpoint e o formato das requisições e respostas.

### `README.md`

Este documento.

---

## ⚙️ Funcionamento da Solução

O fluxo completo pode ser resumido em:

1. **Provisionamento (Terraform)**

   - Cria a instância EC2 na AWS
   - Abre portas necessárias (SSH e HTTP)
   - Exporta o IP da instância para os próximos estágios

2. **Configuração (Ansible)**

   - Instala Docker e dependências
   - Configura o K3s como cluster Kubernetes na instância
   - Garante que o ambiente esteja pronto para receber workloads

3. **Aplicação (Go)**

   - Expõe um endpoint `/analyze` via HTTP POST
   - Recebe um JSON com uma frase e retorna estatísticas da mesma

   Exemplo:

   ```json
   { "sentence": "Hello World from Maputo" }
   ```

   Resposta:

   ```json
   {
     "words": 4,
     "vowels": 8,
     "consonants": 13
   }
   ```

4. **Containerização (Docker)**

   - O código Go é compilado em uma imagem minimalista
   - A imagem é enviada para o **Docker Hub**

5. **Deploy (K3s + Kubernetes)**

   - O cluster aplica os manifests (`deployment.yaml`, `service.yaml`)
   - O serviço fica acessível externamente via NodePort

6. **Pipeline CI/CD (GitHub Actions)**
   - Dispara a cada `git push`
   - Executa automaticamente todas as etapas:
     - `terraform apply` → cria a infra
     - `ansible-playbook` → configura o EC2
     - `docker build` e `docker push` → gera e envia a imagem
     - `kubectl apply` → faz o deploy no K3s
     - Testa o endpoint com `curl` para validar a resposta

---

## 🔍 Evidências de Deploy

- **Pods em execução:**

```bash
kubectl get pods
NAME                          READY   STATUS    RESTARTS   AGE
go-microservice-6d77f7bb9c    1/1     Running   0          3m
```

- **Serviço exposto:**

```bash
kubectl get svc
NAME              TYPE       CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
go-microservice   NodePort   10.43.87.129   <none>        8080:30080/TCP   3m
```

- **Teste da API:**

```bash
curl -X POST http://<EC2-IP>:<NodePort>/analyze      -H "Content-Type: application/json"      -d '{"sentence":"Hello from Mozambique"}'
```

Resposta esperada:

```json
{
  "words": 3,
  "vowels": 9,
  "consonants": 12
}
```

---

## 🌟 Melhorias Futuras

- ✅ Adicionar testes unitários mais completos no Go
- 🔍 Integrar **SonarQube** para análise estática de código
- 📊 Incluir **monitoramento** com Prometheus e Grafana

---

## 📚 Conclusão

Este projeto entrega uma solução **end-to-end** para demonstrar habilidades práticas em **Platform Engineering**.  
O foco não está apenas no código da aplicação, mas na **automação completa do ciclo de vida**:

**infraestrutura → configuração → container → deploy → testes**

Assim, cada alteração no código resulta em uma entrega confiável, reproduzível e pronta para produção.
