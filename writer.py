import data_pb2

def main():

    # We create an instance of the message type "Demo"...
    data = data_pb2.Demo()
    
    # ...and fill it with data
    data.A = long(5)
    data.B = long(5)
    data.C = long(2015)
    
    
    print "* Python writing to file"
    f = open('tst.bin', 'wb')
    f.write(data.SerializeToString())
    f.close()
    
    f = open('tst.bin', 'rb')
    read = data_pb2.Demo()
    read.ParseFromString(f.read())
    f.close()
    
    print "* Python reading from file"
    print "\tDemo.A: %d, Demo.B: %d, Demo.C: %d" %(read.A, read.B, read.C)
    
if __name__ == '__main__':
    main()
    
