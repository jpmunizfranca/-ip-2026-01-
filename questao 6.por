programa
{
    funcao inicio()
    {
        inteiro n, i
        real fahrenheit, celsius

        
        escreva("Digite o número de temperaturas: ")
        leia(n)

        
        para(i = 1; i <= n; i++){
      
        escreva("Digite a temperatura em Fahrenheit: ")
        leia(fahrenheit)

        celsius = 5.0 * (fahrenheit - 32.0) / 9.0

        escreva(fahrenheit, " FAHRENHEIT EQUIVALE A ", celsius, " CELSIUS\n")}
        
        
    }
}
