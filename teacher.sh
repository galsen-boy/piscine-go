$MAIN_SUSPECT="Dartey Henv"
export KEY_INTERVIEW_NUMBER=$(head -n 179 ./streets/Buckingham_Place | tail -n 1 | cut -d '#' -f 2)
echo $KEY_INTERVIEW_NUMBER
cat ./interviews/interview-$KEY_INTERVIEW_NUMBER
echo $MAIN_SUSPECT