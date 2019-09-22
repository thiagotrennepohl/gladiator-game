# Gladiator game

This is a turn based gladiator style game

## Classes
Currently the game has 3 classes

- Mage
- Warrior
- Archer

## Attribute points

Before the game begins (at the character creation screen) you'll be able to set 5 attribute points in any of the following attributes:

- Strenght
- Dextery
- Stamina
- Intelligence
- Luck

Whenever you level up you'll gain 2 more point to add to your attributes.

## Atributos de batalha

- Pontos de vida: 100 pontos
- Ataque: (forca/10 + dano da arma) - defesa-do-oponente
- Defesa: (stamina/8 + valor de armadura(se houver))
- Evasão: (Agilidade/7 + valor de items(se houver))
- Chance de dano crítico: (sorte / 6)%
- Chance de Fuga: (sorte / 6)%

## Oponentes

O jogo criará randomicamente um oponente para o jogador. A cada turno que o jogador vencer, o jogo ficará mais difícil.
Ou seja, o jogo terá que definir randomicamente uma classe (dentro das classes já definidas)

## Inventário

O jogador contará com um inventário de 2 espaços, onde ele poderá carregar equipamentos que o ajudarão na jornada.
Ao fim de cada turno, o jogador será presenteado com um novo item, que será randomicamente escolhido dentro de uma lista de items.
Se o inventário estiver cheio, o jogador deverá escolher se um item do inventário deve ser removido ou se ele rejeitará o item que lhe foi presenteado.

## Batalha

A batalha é baseada em turnos e o jogador poderá fazer apenas um movimento por vez.
Durante a batalha, o jogador poderá escolher qual ação seu personagem poderá fazer. Essas ações incluem:

- Atacar
- Defender
- Usar item
- Fugir

Ao fim de cada batalha em que o jogador for vitorioso, o jogador poderá atribuir 1 novo ponto em um dos atributos de seu personagem.

# TODO

- Adicionar armadilhas
- Jogo online

# ENGLISH - Medieval Battle

...
