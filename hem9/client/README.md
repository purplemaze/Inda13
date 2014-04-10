FIXME
I've found two insidious bugs in this function; both of them are unlikely
to show up in testing. Please fix them right away – and don't forget to
write a doc comment this time.

The first bug is a data race which occurs when both the gorutine and the "timeout case" writes to the res var.
The second bug occurs when "the timeout case" is executed before the gorutine is done. 
The gorutine then blocks forever because "no one" is receiving. 
This creates a memory leak and if the program runs long enough it will crash.
Both bugs can be fixed by creating a new buffered channel.
I got inspired and helped by: https://vividcortex.com/blog/2014/01/15/two-go-memory-leaks/
