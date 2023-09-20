# Jogo do Bicho API

É um projeto onde vamos criar uma API utilizando Golang como tecnologia principal da aplicação. 

Ferramentas da aplicação: 

- Tecnologia:
    - Golang
- Banco de dados:
    - MongoDB (Bet Data)
    - Postgress (User data)
- Cache:
    - Redis
- Arquitetura:
    - Hexagonal

**************************************Fluxo da aplicação:**************************************

***Backend{user}:***

- Recebe os dados do usuários
- Envia para o banco de dados
- Login (com autenticação)
- Novo usuário, recebe 15 moedas para começar a jogar
- Cada aposta só pode ser de no minimo 1 animal e no maximo 3 animais.
- Vai ter um quantidade 15 de moedas
    - Todo usuario ao entrar na plataforma, recebe um valor de moedas para poder comprar tickets
    - Cada ticket custa 5 moedas
    - Cada vitoria multiplica por 3x o valor de moedas apostado
    - Se perder o dinheiro vai para a banca
- Se perder todo o dinheiro, pode “pagar” para ter mais moedas.
- Cada  5 moeda no jogo, custa 10,00R$
- Para sacar, só pode se tiver mais de 150 moedas
- Valor de 5 moedas para saque, custa 5,00 R$

***Backend{sorteio}:***

- Sorteia-se 5 numeros
- Numeros sempre em pares. Ex: 01, 20,55
- Os numeros são de 0 ~ 99
- E cada numero representa um animal
- Os usuários campeões recebem a informação de que foi campeão
- O sorteio é realizado ao bater na rota (”/sorteio”)

***Backend{bet}:***

- **Aposta Simples**: Você pode escolher vários números e fazer apostas simples em cada um deles. Por exemplo, você pode escolher os números 7 (Carneiro), 11 (Cavalo) e 18 (Porco) e fazer uma aposta em cada um deles separadamente. Se algum desses números for sorteado, você ganha com base na aposta feita para esse número específico.
- **Grupos**: Em vez de fazer várias apostas simples, você pode optar por apostar em um grupo de números. Os grupos são combinações de números que têm probabilidades diferentes de ganhar e pagam prêmios diferentes. Por exemplo, você pode fazer uma aposta em um "grupo de 4" que inclui os números 7, 11, 18 e 25. Se qualquer um desses números for sorteado, você ganha de acordo com as regras do grupo.

***Backend{animais}:***

- Cada animal, representa uma dessas casas:
1. Avestruz
2. Águia
3. Burro
4. Borboleta
5. Cachorro
6. Cabra
7. Carneiro
8. Camelo
9. Cobra
10. Coelho
11. Cavalo
12. Elefante
13. Galo
14. Gato
15. Jacaré
16. Leão
17. Macaco
18. Porco
19. Pavão
20. Peru
21. Touro
22. Tigre
23. Urso
24. Veado
25. Vaca

***Obs: Criar regra, para “puxar” do banco de dados, referente a cada numero.*** 

***User:***

- Usuario se cadastra na plataforma
- Usuario pode comprar ‘N’ tickets
- Se sorteado, o mesmo recebe 3x o valor da quantidades de tickets que ele recebeu.