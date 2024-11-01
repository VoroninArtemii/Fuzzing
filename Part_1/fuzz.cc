#include "pugixml/src/pugixml.hpp"
#include <iostream>
#include <chrono>

using namespace std;

extern "C" int LLVMFuzzerTestOneInput(const uint8_t *Data, size_t Size){

/*
//	The output of mutated data in console
	for (int i = 0; i < Size; i++)
		cout<<Data[i];
	cout<<endl<<endl<<endl<<endl;
*/

/*
//	Current time
	auto end = std::chrono::system_clock::now();
	std::time_t end_time = std::chrono::system_clock::to_time_t(end);
	std::cout<<std::ctime(&end_time)<< std::endl;
*/

	pugi::xml_document doc;
	doc.load_buffer(Data, Size);
	return 0;
}
