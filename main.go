package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World");
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite um número: ")
	input, _ := reader.ReadString('\n')

	var numero int
	fmt.Sscanf(input, "%d", &numero)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", numero, i, numero*i)
	}
}
// ------------------------------------------------- VAGA ---------------------------------------------------------------------

// Descrição do trabalho
// Já conhece a GFT?

// Com mais de 35 anos de mercado e localizada em 16 países, a GFT emprega mais de 10.000 pessoas ao redor do mundo sendo 3.000 pessoas no Brasil.

// Aqui, a diversidade, a igualdade e a inclusão estão no centro e garantir um ambiente de trabalho seguro para todas as pessoas faz parte de quem somos.

// Promovemos uma cultura de desenvolvimento e crescimento pautadas no nosso pilar de Continuous Scaled Learning porque acreditamos o uso inteligente da tecnologia é a chave para o sucesso nesse mundo digital.

// O que buscamos

// Profissional de nível Sênior que atue com Golang e AWS.

// Requisitos e Qualificações
// • Experiência sólida com Golang (nível avançado)
// • Experiência com microsserviços e sistemas distribuídos
// • EC2, S3, RDS, DynamoDB
// • Lambda, API Gateway
// • SNS, SQS, EventBridge
// • Práticas de observabilidade
// • Experiência com Docker e orquestração (Kubernetes desejável)
// • Forte conhecimento em APIs REST e/ou gRPC
// • Experiência com bancos relacionais e NoSQL
// • Vivência com CI/CD (GitHub Actions)
// • Conhecimento em segurança de aplicações e boas práticas cloud
// • Experiência com Git (versionamento de código)
// • Vivência em metodologias ágeis
// Descrição comportamental:
// Procuramos uma pessoa que
// • Goste de trabalhar em equipe e seja colaborativa em suas atribuições;
// • Tenha coragem para se desafiar e ir além, abraçando novas oportunidades de crescimento;
// • Transforme ideias em soluções criativas e busque qualidade em toda sua rotina;
// • Tenha habilidades de resolução de problemas ;
// • Possua habilidade e se sinta confortável para trabalhar de forma independente e gerenciar o próprio tempo ;
// • Tenha interesse em lidar com situações adversas e inovadoras no âmbito tecnológico.
// O que oferecemos:
// • Cartão multi-benefícios – você escolhe como e onde utilizar;
// • Bolsas de Estudos para cursos de Graduação, Pós, MBA e Idiomas;
// • Programas de incentivo à Certificações;
// • Horário de trabalho flexível;
// • Plano personalizado de carreira e possibilidade de carreira internacional;
// • Salários competitivos;
// • Avaliação de desempenho anual com possibilidade de bonificação;
// • Gympass;
// • Previdência Privada;
// • Auxílio-Creche;
// • Assistência Médica;
// • Assistência Odontológica;
// • Seguro de Vida.
// Big enough to deliver – small enough to care.



// ------------------------------------------------- FAIXA SALARIAL ---------------------------------------------------------------------
// Júnior: R$ 4.000 até R$ 7.000; 
// Pleno:  R$ 7.000 até R$ 12.000; 
// Sênior: R$ 12.000 até R$ 18.000.



// ------------------------------------------------- PARADIGMA ---------------------------------------------------------------------

// MULTIPARADIGMA: Procedural, Orientado a objetos, e concorrente.
