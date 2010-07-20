TARGET = 6.out
OBJS = glassDice.6

$(TARGET): $(OBJS)
	6l -o $@ $^

%.6: %.go
	6g -o $@ $<
