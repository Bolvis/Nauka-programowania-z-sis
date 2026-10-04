#include <iostream>

const int DAWID_WIEK = 25;

int main()
{
    std::string imię;
    std::cout << "Jak masz na imię?" << std::endl;
    std::cin >> imię;
    std::cout << "Cześć " << imię << "!" << std::endl;

    std::string wiek;
    std::cout << "A ile masz lat?" << std::endl;
    std::cin >> wiek;
    std::cout << "Super, to znaczy, że ja mam " << DAWID_WIEK << " a ty " << imię << " masz " << wiek << std::endl;

    return 0;
}