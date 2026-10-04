#include <iostream>
#include <string>

const int DAWID_WIEK = 25;

bool        czy_autor_młodszy(int wiek_użytkownika);
bool        czy_autor_w_tym_samym_wieku(int wiek_użytkownika);
int         różnica_wieku(int wiek_użytkownika, bool autor_młodszy);
std::string odmiana_gramatyczna_lat(int różnica_wieku);

int main()
{
    std::string imię;
    std::cout << "Jak masz na imię?" << std::endl;
    std::cin >> imię;
    std::cout << "Cześć " << imię << "!" << std::endl;

    int wiek;
    std::string wiek_tekst;
    std::cout << "A ile masz lat?" << std::endl;
    std::cin >> wiek_tekst;
    wiek = std::stoi(wiek_tekst);

    bool autor_młodszy = czy_autor_młodszy(wiek);
    bool autor_w_tym_samym_wieku = czy_autor_w_tym_samym_wieku(wiek);
    if (autor_w_tym_samym_wieku)
    {
        std::cout << "Super, jesteśmy w tym samym wieku :)" << std::endl;
    } else
    {
        int różnica = różnica_wieku(wiek, autor_młodszy);
        std::cout << "WOW, to mam "
            <<  std::to_string(różnica)
            << odmiana_gramatyczna_lat(różnica)
            << (autor_młodszy ? "mniej" : "więcej") << std::endl;
    }

    return 0;
}

bool czy_autor_młodszy(int wiek_użytkownika)
{
    return DAWID_WIEK < wiek_użytkownika;
}

bool czy_autor_w_tym_samym_wieku(int wiek_użytkownika)
{
    return DAWID_WIEK == wiek_użytkownika;
}

int różnica_wieku (int wiek_użytkownika, bool autor_młodszy)
{
    return autor_młodszy ? wiek_użytkownika - DAWID_WIEK : DAWID_WIEK - wiek_użytkownika;
}

std::string odmiana_gramatyczna_lat(int różnica)
{
    if (różnica == 1) {
        return " rok ";
    }

    int ostatnia_cyfra = różnica % 10;
    int nast_ostatnia = (różnica / 10) % 10;

    if (nast_ostatnia != 1 && (ostatnia_cyfra >= 2 && ostatnia_cyfra <= 4)) {
        return " lata ";
    }

    return " lat ";
}