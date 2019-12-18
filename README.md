# loremfile :scroll:

Quick CLI utility to generate lorem ipsum ish text files of a specific size.
Useful for benchmarking "wordcount" type programs.

The random seed used will be printed to stderr. You can then use the `-seed`
argument to reproduce the exact same text file on another system. This is useful
to not have to distribute a large multi-gigabyte test file to others.

## Usage

### Help
```shell
$ loremfile -help
Usage of loremfile:
  -seed int
        use specific random seed
  -size string
        desired size of output (default "32kb")
```

### Sample
```shell
$ loremfile -size 1kb > test1.txt
Generating >= 1024 bytes of lorem ipsum with seed 1576681389603936000

$ fmt test1.txt 
Stipendium teneam si possemus excellentiam nam manducantem gradibus
quam audire hac elapsum re pro aut agam cordis violari pro dicimur.
Amo pro re, tota dubia.

Vivant natura docens abs te naturae es cura me os turpibus ea.
Frigidique tolerat habitaculum libeat ipse semper nam nominum dicant
fui gaudent deo doctrinae claudicans malum quantulum expertum ponere.
Me exemplo sensibus horum eas aula cor, indicatae exclamaverunt
teneo me cur ne cadere conduntur ore peccato abditis.

Me nosse o te rogo sunt satietas re requirunt. Qui surgere ieiuniis
ne agro a divino oblitum ex hierusalem da se. Hac loqueremur deum
laudare habeas manducantem longum ibi cor certissimum contristentur
hi nutu deum noe hi e a nota a bellum. Meridies peccare loqueremur
latet ore. Exclamaverunt vi catervas habemus occideris latet calamitas
at mei homo meam des cognosceremus religione caro indica pro inpiorum
nollent dulcedo miris.

Fleo plenariam me vi ceterorum te sapere cognitor teneo dura cor
deum misericors lucentem dedisti, oportet faciant benedicere rapit.

$ ./loremfile -size 1kb -seed 1576681389603936000 > test2.txt
Generating >= 1024 bytes of lorem ipsum with seed 1576681389603936000

$ md5sum test*.txt
90623ffcff258d705f58bdd690c24442  test1.txt
90623ffcff258d705f58bdd690c24442  test2.txt
```