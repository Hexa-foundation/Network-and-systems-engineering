//Author Andrew mbugua

/*is used to prep
the socket address structures for subsequent use*/

#include <iostream>
#include <unistd.h>
#include <sys/socket.h>
using namespace std;

struct addrinfo{
	
	  int ai_flags;
	  int ai_family;
    int ai_sockettype;
    int ai_protocol;
    size_t ai_addrlen;
    
    struct sockaddr *ai_addr;
	  char *ai_canonname; //full canonical host name
	
	struct addrinfo *ai_next; //linked list,next node
};


int main()
{

	
	return 0;
}
