programa
{
    inclua biblioteca Matematica --> mat

    funcao inicio()
    {
        real h, a, Ab, V

        escreva("Digite a altura e a aresta.")
        leia(h, a)

        Ab = (3.0 * a * a * mat.raiz(3.0, 2.0)) / 2.0

        V = (1.0 / 3.0) * Ab * h
        
        V = mat.arredondar(V, 2)

        escreva("O VOLUME DA PIRAMIDE E = ", V, " METROS CUBICOS\n")
    }
}
