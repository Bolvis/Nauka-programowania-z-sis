#include <iostream>

int main()
{
    std::string imię;
    std::cout << "Jak masz na imię?" << std::endl;
    std::cin >> imię;
    std::cout << "Cześć " << imię << "!" << std::endl;

    return 0;
}