#include "pugixml/src/pugixml.hpp"
#include <stdio.h>

extern "C" int LLVMFuzzerTestOneInput(const uint8_t *Data, size_t Size){
	pugi::xml_document doc;
//	for (int i = 0; i < Size; i++)
//		printf("%c", Data[i]);
//	printf("\n");
	doc.load_buffer(Data, Size);
	return 0;
}
