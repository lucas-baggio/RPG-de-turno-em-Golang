package main

import (
	"fmt"
	"time"
)

type Player struct {
	Nome   string
	Vida   int
	Ataque int
	Defesa int
	Level  int
}

func main() {
	jogador := Player{"Lucas", 100, 20, 10, 1}

	fmt.Println("Bem-vindo ao Coliseu!")
	fmt.Println("Você é um jovem guerreiro em busca de glória.")
	fmt.Print("Deseja começar a batalhar? (s/n): ")

	var resposta string
	fmt.Scan(&resposta)

	if resposta != "s" {
		fmt.Println("Você decidiu não lutar hoje. O Coliseu aguarda seu retorno...")
		return
	}

	fmt.Println("\nO primeiro inimigo se aproxima!")
	batalhar(&jogador, &Player{"Goblin", 60, 15, 5, 1})

	if jogador.Vida > 0 {
		subirDeNivel(&jogador)
		fmt.Println("\n Um novo desafio o aguarda...")
		time.Sleep(2 * time.Second)
		batalhar(&jogador, &Player{"Orc Selvagem", 90, 18, 8, 2})
	}

	if jogador.Vida > 0 {
		subirDeNivel(&jogador)
		fmt.Println("\nVocê conquistou o Coliseu! Todos aclamam seu nome!")
	}
}

func batalhar(jogador, inimigo *Player) {
	fmt.Printf("\n%s (Lv %d) vs %s (Lv %d)\n", jogador.Nome, jogador.Level, inimigo.Nome, inimigo.Level)
	fmt.Println("=====================================")

	for jogador.Vida > 0 && inimigo.Vida > 0 {
		fmt.Printf("\n%s: %d HP | %s: %d HP\n", jogador.Nome, jogador.Vida, inimigo.Nome, inimigo.Vida)
		fmt.Println("1 - Atacar")
		fmt.Println("2 - Defender")
		fmt.Println("3 - Poção")
		fmt.Print("> ")

		var escolha int
		fmt.Scan(&escolha)

		switch escolha {
		case 1:
			atacar(jogador, inimigo)
		case 2:
			defender(jogador)
		case 3:
			usarPocao(jogador)
		default:
			fmt.Println("Opção inválida.")
			continue
		}

		if inimigo.Vida > 0 {
			time.Sleep(1 * time.Second)
			contraAtaque(inimigo, jogador)
		}
		time.Sleep(1 * time.Second)
	}

	if jogador.Vida <= 0 {
		fmt.Println("\n Você foi derrotado no Coliseu.")
	} else {
		fmt.Printf("\n%s derrotou %s!\n", jogador.Nome, inimigo.Nome)
	}
}

func atacar(a, b *Player) {
	dano := a.Ataque - b.Defesa/2
	if dano < 0 {
		dano = 0
	}
	b.Vida -= dano
	fmt.Printf("%s ataca %s causando %d de dano!\n", a.Nome, b.Nome, dano)
}

func defender(p *Player) {
	p.Defesa += 3
	fmt.Printf("%s se defende, aumentando a defesa!\n", p.Nome)
}

func usarPocao(p *Player) {
	cura := 20
	p.Vida += cura
	fmt.Printf("%s usou uma poção e recuperou %d de vida!\n", p.Nome, cura)
}

func contraAtaque(a, b *Player) {
	dano := a.Ataque - b.Defesa/3
	if dano < 0 {
		dano = 0
	}
	b.Vida -= dano
	fmt.Printf("%s contra-ataca e causa %d de dano!\n", a.Nome, dano)
}

func subirDeNivel(p *Player) {
	p.Level++
	p.Vida += 30
	p.Ataque += 5
	p.Defesa += 3
	fmt.Printf("\n %s subiu para o nível %d!\n", p.Nome, p.Level)
	fmt.Printf("Seus atributos aumentaram!\nVida: %d | Ataque: %d | Defesa: %d\n", p.Vida, p.Ataque, p.Defesa)
}
