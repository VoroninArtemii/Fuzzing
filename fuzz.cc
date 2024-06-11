#include "pugixml/src/pugixml.hpp"
#include <iostream>

using namespace std;

extern "C" int LLVMFuzzerTestOneInput(const uint8_t *Data, size_t Size){

/*
//	The output of mutated data
	for (int i = 0; i < Size; i++)
		cout<<Data[i];
	cout<<endl<<endl<<endl<<endl;
*/

	pugi::xml_document doc;
	doc.load_buffer(Data, Size);
	return 0;
}
