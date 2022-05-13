#! /bin/bash
TAGS=`git describe --tags --always`
B1=`echo $TAGS | awk '{split($0,a,"."); print a[1]}'`
B2=`echo $TAGS | awk '{split($0,a,"."); print a[2]}'`
B3=`echo $TAGS | awk '{split($0,a,"."); print a[3]}'`
B3=`echo $B3 | awk '{split($0,a,"-"); print a[1]}'`
#echo $TAGS "---" B1 $B1 "---" B2 $B2 "---" B3 $B3 --------$B4
Yeni=${B1}.${B2}.$((B3+1))
echo $Yeni
git tag $Yeni
git push --tags