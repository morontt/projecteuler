# Largest exponential
#
# Completed on Thu, 30 Jul 2015, 20:03

sed 's/,/\ /' p099_base_exp.txt | awk '{print log($1)*$2 " " NR}' | sort | tail -n 1
