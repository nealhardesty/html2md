html 2 markdown converter

uses golang

must have a Makefile for build and test

uses github.com/JohannesKaufmann/html-to-markdown as the converter internally

usage: html2md <inputfile> [outputfile] 

if no outputfile, just output to stdout

if no inputfile, assume stdin
