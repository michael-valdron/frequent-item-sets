#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Created on Mon Jan 22 17:58:20 2018

@author: Michael Valdron
"""

import sys
import apriori

fname = "../../data/test.dat"

def main(args, argc):
    f = open(fname)

    items = apriori.get_unique_items(f.readlines())
    
    f.close()
    
    print(items)
    
    return 0


if __name__ == "__main__":
    sys.exit(main(sys.argv, len(sys.argv)))