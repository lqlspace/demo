grammar ApiParser;

import ApiLexer;

api:            spec*;
spec:           typeSpec
                |serviceSpec
                ;

typeSpec:       typeToken='type' ID structToken='struct' lbrace='{'  field* rbrace='}';
field:          ID ID tag=RAW_STRING?;

serviceSpec:    ID serviceName lbrace='{' serviceRoute* rbrace='}';
serviceRoute:   atHandler route;
atHandler:      ATHANDLER ID;
route:          ID path request? response?;
request:        lp='(' (ID)? rp=')';
response:      returnToken='returns' lp='(' ID? rp=')';
serviceName:    (ID '-'?)+;
path:           (('/' (ID ('-' ID)*))|('/:' (ID ('-' ID)?)))+ | '/';