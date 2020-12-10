# tap
A parallel architecture trace parser using the PAT trace format.

# compile
make build

# run example
make run

# pat 
The PAT trace format is composed of a number of tokens, which are as follows:

**Dimensions:**
*JD* jx jy jz wx wy wz

**Thread ID:**
*T* tid

**Instruction Packet:**
*P* pc #instructions #dependent_addrs dependent_addrs*

**Load Packet:**
*L* #bytes-1 address pc #instructions #dependent_addrs dependent_addrs*

**Store Packet:**
*S* #bytes-1 address pc #instructions #dependent_addrs dependent_addrs*

**Barrier Packet:**
*B* pc #instructions #dependent_addrs dependent_addrs*

