Based on the docs of the DCOM Protocol (https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dcom) 
and the implementation of the aiowmi library (https://github.com/cesbit/aiowmi)
I will start implementing necessary the Common Data Types (See MS-DCOM, 2.2)
These are considered the building blocks of the DCOM Protocol.

The protocol is built on top of the RPC specifications and Aiowmi using NDR (Network Data Representation) to transmit the data in a platform-independent manner

 Working on right now:
 - ORPCTHAT and ORPCTHIS data structures: it's used for sending additional information in the RPC call in a platform-independent manner
 
 - OBJREF: it's the format in use to marshal the DCOM Object Reference, there are four types of OBJREF, will implement only two for now OBJREF_STANDARD and OBJREF_CUSTOM
 
 - DUAL_STRING_ARRAY
