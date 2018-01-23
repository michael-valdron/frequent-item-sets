#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Created on Mon Jan 22 17:58:20 2018

@author: Michael Valdron
"""

def get_unique_items(fcontents):
    items = []
    for b in fcontents:
        for item in str(b).strip().split(' '):
            if item != '':
                items.append(int(item))

    return items
